uint[] memory a = new uint[](7); // uintで5つの動的配列の配列を宣言
a[6] = 8; // 動的配列を書き換え
bytes memory b = new bytes(5); // 一時的な動的配列を宣言
b[0] = '0'; // 動的配列bを書き換え
string memory s = 'keyword';
// s[0] = '1'; // stringにはindexでアクセスできない
