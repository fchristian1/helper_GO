package helper_GO

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"math/big"
	"os"
	"path/filepath"
	"time"
)

func InitTLS(sslPath, sslCert, sslKey string, logLevel string) {
	//if tls files not found create new ones
	_, certErr := os.Stat(filepath.Join(sslPath, sslCert))
	_, keyErr := os.Stat(filepath.Join(sslPath, sslKey))

	// If TLS files not found, create new ones
	if os.IsNotExist(certErr) || os.IsNotExist(keyErr) {
		// Generate new TLS files
		cert, key, err := generateTLSFiles()
		if err != nil {
			LogError(logLevel, "Failed to generate TLS files:", err)
		}
		//check file path and create if not exists
		err = os.MkdirAll(sslPath, 0755)
		if err != nil {
			LogError(logLevel, "Failed to create directory:", err)
		}
		// Write the generated files to disk
		err = os.WriteFile(filepath.Join(sslPath, sslCert), cert, 0644)
		if err != nil {
			LogError(logLevel, "Failed to write TLS certificate file:", err)
		}
		err = os.WriteFile(filepath.Join(sslPath, sslKey), key, 0644)
		if err != nil {
			LogError(logLevel, "Failed to write TLS key file:", err)
		}
		LogInfo(logLevel, "TLS files are generated in: ")
	}
	LogInfo(logLevel, "TLS files are found in: ")
}

// generation of a crt and a key file for TLS with given data
func generateTLSFiles() ([]byte, []byte, error) {
	// Generate a new private key
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	// Encode the private key to PEM
	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	// Generate a new self-signed certificate
	cert, err := generateSelfSignedCertificate(key)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate self-signed certificate: %w", err)
	}

	// Encode the certificate to PEM
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})

	return certPEM, keyPEM, nil
}

// Generate a self-signed certificate
func generateSelfSignedCertificate(key *rsa.PrivateKey) ([]byte, error) {
	// Create a certificate template
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 365 * 10),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Generate the certificate
	return x509.CreateCertificate(nil, template, template, &key.PublicKey, key)
}
