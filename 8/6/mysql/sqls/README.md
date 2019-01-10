# SQLファイル

## 概要

*.sqlファイルが辞書順に自動実行されます。

### 001_initialize.sql

テーブルの初期化を行います。

### 002_sample-data.sql

サンプルデータの投入を行います。

## dummy-data.rb

サンプルデータの自動生成を行います。

### 実行手順

```
$ sudo apt install ruby
$ sudo gem install faker
$ sudo gem install digest/sha2
$ ./dummy-data.rb
```

### 設定

サンプルデータの生成個数は、dummy-data.rb内の下記箇所を修正してください。

```
USER_NUM = 10
ROOM_NUM = 20
QUESTION_NUM = 40
```

## XXX_describe-select-statements.sql-test

チューニング用のサンプルクエリです。

### 実行手順例

```
$ mysql -h 127.0.0.1 -u root -p < XXX_describe-select-statements.sql-test -t
```
