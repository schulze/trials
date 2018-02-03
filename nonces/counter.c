#include <stdio.h>

// print prints the first 8 bytes of nonce.
void print(unsigned char nonce[24]) {
	for (int i = 0; i < 8; i++) {
		printf("%lu ", (unsigned long int)nonce[i]);
	}
	printf("\n");
}

void increment_nonce(unsigned char nonce[24]) {
	unsigned int s = 1;
	for (int i = 0; i < 24; i++) {
		s += (unsigned int)nonce[i];
		nonce[i] = (unsigned char)s;
		s >>= 8;
	}
}

int main() {
	unsigned char nonce[24];
	for (int i = 0; i < 24; i++) {
		nonce[i] = 0;
	}
	for (int i = 0; i < 1000; i++) {
		print(nonce);
		increment_nonce(nonce);
	}
	return 0;
}

