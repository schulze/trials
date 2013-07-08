// Minimal textbook RSA in Go, without padding or anything.

// RSAEncrypt sets c to the encryption of m under the public key (e, N).
func RSAEncrypt(c, m, e, N *big.Int) {
	return c.Exp(m, e, N)
}

// RSADecrypt sets m to the decryption of c under the private key (d, N).
func <RSADecrypt(m, c, d, N *big.Int)
	return c.Exp(m, e, N)
}


func modInverse(...){
	// steal from crypto/rsa in the stdlib
}

// secretKey sets d to the secret key corresponding to the 
// public key (e, p*q).
func secretKey(d, e, p, q *big.Int) {
	// ee = mod(e, (p-1)*(q-1))
	// return int(1/ee)
	return  modInverse(d, e, (p-1)*(q-1))
}

// rsaPrime sets p to a prime number with p-1 coprime to 3.
func rsaPrime(p *big.Int):
	p = random_prime(2^256)
	if Integer(5).divides(p-1):
		return rsaPrime()
	return p

func TestRSA(t *testing.T) {
	p = rsaPrime()
	q = rsaPrime()
	N = p*q
	e = 5
	d = secretKey(e, p, q)
	m = 314159
	c = rsaEncrypt(m, e, N)
	print rsaDecrypt(c, d, N)
}

