#!/usr/bin/ruby
require 'faker'
require 'digest/sha2'
require 'set'

# ===========================
#  サンプルデータの挿入SQL生成ツール
# ===========================

USER_NUM = 10 # usersのレコード数
ROOM_NUM = 20 # roomsのレコード数
QUESTION_NUM = 40 # questionsのレコード数
RECORDS_PER_SENTENCE = 100 # 一つのINSERT文で挿入するレコード数
OUTPUT_FILE = '002_sample-data.sql' # 出力ファイル名

Faker::Config.locale = :ja # サンプルデータの言語（デフォルトは日本語）


# ===========================
#  前処理
# ===========================

file = File.open(OUTPUT_FILE, "w")

# ===========================
#  Usersデータの生成
# ===========================
User = Struct.new(:id, :ethaddress, :nickname, :mail, :encrypted_password, :salt)
users = []

for num in 1..USER_NUM do
  # ===========================
  # Userデータの生成
  # ===========================
  # saltをランダムな20文字の文字列で生成
  salt = Faker::Internet.password(20)
  # パスワードは固定で "password"
  password = "password"
  # sha256でパスワードをハッシュ化
  encrypted_password = Digest::SHA256.hexdigest(password + salt)
  user = User.new(num, Faker::Ethereum.unique.address, Faker::Name.name, Faker::Internet.unique.email, encrypted_password, salt)
  users << user
  
  # ===========================
  # SQLの出力
  # ===========================
  # RECORDS_PER_SENTENCEごとにINSERT句を出力
  if num % RECORDS_PER_SENTENCE == 1
    file.puts("INSERT INTO `users`(`id`, `ethaddress`, `nickname`, `mail`, `encrypted_password`, `salt`) VALUES")
  end
  # 最後のデータかRECORDS_PER_SENTENCEの最後の場合はセミコロン、それ以外はカンマを使用
  separator = ((num == USER_NUM) or (num % RECORDS_PER_SENTENCE == 0)) ? ";" : ","
  # データ部分の出力
  file.puts("  (%d, '%s', '%s', '%s', '%s', '%s')%s" % [user.id, user.ethaddress, user.nickname, user.mail, user.encrypted_password, user.salt, separator])
  file.flush
end

# ===========================
#  Roomsデータの生成
# ===========================
Room = Struct.new(:id, :owner_id, :owner_address, :title, :event_code, :address, :create_tx_hash)
rooms = []

# len桁のイベントコードを生成する関数
def gen_event_code(len)
  # Base58ベースの数字と大文字からランダムな文字列を生成
  return "23456789ABCDEFGHJKLMNPQRSTUVWXYZ".split("").shuffle.first(len).join
end

# num個のユニークなlen桁のイベントコードを生成する関数
def gen_unique_event_codes(len, num)
  codes = Set.new
  loop do
    codes << gen_event_code(len)
    # イベントコードがnum個になるまで再帰処理
    return codes.to_a if codes.size >= num
  end
end

# 4桁のイベントコードをROOM_NUM個生成する
event_code_set = gen_unique_event_codes(4, ROOM_NUM)

for num in 1..ROOM_NUM do
  # ===========================
  # Roomデータの生成
  # ===========================
  # 関連するユーザの取得
  owner = users[rand(1..USER_NUM) - 1]
  # トランザクションハッシュの生成（本来はkeccak256だがsha256で代用する）
  tx_hash = "0x" + Digest::SHA256.hexdigest(Faker::Internet.password(20))
  room = Room.new(num, owner.id, owner.ethaddress, Faker::App.name, event_code_set[num], Faker::Ethereum.unique.address, tx_hash)
  rooms << room
  # ===========================
  # SQLの出力
  # ===========================
  # RECORDS_PER_SENTENCEごとにINSERT句を出力
  if num % RECORDS_PER_SENTENCE == 1
    file.puts("INSERT INTO `rooms`(`id`, `owner_id`, `owner_address`, `title`, `event_code`, `address`, `create_tx_hash`) VALUES")
  end
  # 最後のデータかRECORDS_PER_SENTENCEの最後の場合はセミコロン、それ以外はカンマを使用
  separator = ((num == ROOM_NUM) or (num % RECORDS_PER_SENTENCE == 0)) ? ";" : ","
  # データ部分の出力
  file.puts("  (%d, %d, '%s', '%s', '%s', '%s', '%s')%s" % [room.id, room.owner_id, room.owner_address, room.title, room.event_code, room.address, room.create_tx_hash, separator])
  file.flush
end

# ===========================
#  Questionsデータの生成
# ===========================
Question = Struct.new(:id, :room_id, :address, :owner_id, :body)
questions = []

for num in 1..QUESTION_NUM do
  # RECORDS_PER_SENTENCEごとにINSERT句を出力
  if num % RECORDS_PER_SENTENCE == 1
    file.puts("INSERT INTO `questions`(`id`, `room_id`, `address`, `owner_id`, `body`) VALUES")
  end
  # 最後のデータかRECORDS_PER_SENTENCEの最後の場合はセミコロン、それ以外はカンマを使用
  separator = ((num == QUESTION_NUM) or (num % RECORDS_PER_SENTENCE == 0)) ? ";" : ","
  # ===========================
  # Questionデータの生成
  # ===========================
  # 関連ルームの取得
  room = rooms[rand(1..ROOM_NUM) -1]
  # 関連ユーザの取得
  owner = users[rand(1..USER_NUM) - 1]
  # ランダムに匿名ユーザのレコードに変換
  if(rand(1..10) >= 5)
    # ログインユーザによるQuestion
    q = Question.new(num, room.id, owner.ethaddress, owner.id, Faker::Lorem.question)
    questions << q
    # データ部分の出力
    file.puts("  (%d, %d, '%s', %d, '%s')%s" % [q.id, q.room_id, q.address, q.owner_id, q.body, separator])
  else
    # 匿名ユーザによるQuestion
    q = Question.new(num, room.id, Faker::Ethereum.unique.address, nil, Faker::Lorem.question)
    questions << q
    # データ部分の出力
    file.puts("  (%d, %d, '%s', NULL, '%s')%s" % [q.id, q.room_id, q.address, q.body, separator])
  end
  file.flush
end

# ===========================
#  後処理
# ===========================
file.close
