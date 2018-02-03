#include <string>
#include <iostream>

using std::string;
using std::cin;
using std::cout;
using std::endl;

// print prints the first 8 bytes of nonce.
void print(string nonce) {
	for (int i = 0; i < 8; i++) {
		cout << (unsigned int)(unsigned char)nonce[i] << " ";
	}
	cout << endl;
}

void increment_nonce(string nonce) {
	unsigned int s = 1;
	for (int i = 0; i < 24; i++) {
		s += (unsigned int)(unsigned char)nonce[i];
		nonce[i] = (unsigned char)(s);
		s >>= 8;
	}
}

int main() {
	string nonce(0, 24);
	for (int i = 0; i < 1000; i++) {
		print(nonce);
		increment_nonce(nonce);
	}
	return 0;
}

