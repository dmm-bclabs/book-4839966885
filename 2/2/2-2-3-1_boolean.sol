bool a = true;
bool b = false;
bool c = a || b; // aがtrueのため、bは評価されず、cはtrueになります
bool d = b && a; // bがfalseのため、aは評価されず、cはfalseになります
