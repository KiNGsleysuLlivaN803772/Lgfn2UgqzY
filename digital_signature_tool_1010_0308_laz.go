// 代码生成时间: 2025-10-10 03:08:29
package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "encoding/asn1"
    "encoding/hex"
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

// Command line flags
var message string
var privateKeyPath string
var publicKeyPath string
var signaturePath string

// asn1Dss1 is the ASN.1 signature structure
type asn1Dss1 struct {
    R, S *big.Int
}

// NewCommand creates a new digital signature tool command
func NewCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "digital-signature-tool",
        Short: "A tool for signing messages with ECDSA",
        Long:  `This tool allows you to sign messages using ECDSA and verify signatures.`,
        Run:   run,
    }
    cmd.Flags().StringVarP(&message, "message", "m", "", 
        "The message to be signed")
    cmd.Flags().StringVarP(&privateKeyPath, "private-key", "p", "", 
        "Path to the private key file")
    cmd.Flags().StringVarP(&publicKeyPath, "public-key", "k", "", 
        "Path to the public key file")
    cmd.Flags().StringVarP(&signaturePath, "signature", "s", "", 
        "Path to the output signature file")
    return cmd
}

// run is the main execution function for the command
func run(cmd *cobra.Command, args []string) {
    // Check if all required flags are provided
    if message == "" || privateKeyPath == "" || publicKeyPath == "" || signaturePath == "" {
        fmt.Println("Error: All flags are required.")
        return
    }

    // Load private key
    privateKey, err := os.ReadFile(privateKeyPath)
    if err != nil {
        fmt.Printf("Error reading private key: %s
", err)
        return
    }

    // Load public key
    publicKey, err := os.ReadFile(publicKeyPath)
    if err != nil {
        fmt.Printf("Error reading public key: %s
", err)
        return
    }

    // Sign the message
    hash := sha256.Sum256([]byte(message))
    signature, err := sign(hash[:], privateKey)
    if err != nil {
        fmt.Printf("Error signing message: %s
", err)
        return
    }

    // Write signature to file
    err = os.WriteFile(signaturePath, signature, 0644)
    if err != nil {
        fmt.Printf("Error writing signature: %s
", err)
        return
    }

    fmt.Println("Message signed successfully.")
}

// sign creates an ECDSA signature of the provided hash
func sign(hash []byte, privateKey []byte) ([]byte, error) {
    // Unmarshal the private key
    privKey, err := x509.ParseECPrivateKey(privateKey)
    if err != nil {
        return nil, err
    }

    // Create the signature
    r, s, err := ecdsa.Sign(rand.Reader, privKey, hash)
    if err != nil {
        return nil, err
    }

    // Marshal the signature into ASN.1
    signature := asn1Dss1{r, s}
    asn1Bytes, err := asn1.Marshal(signature)
    if err != nil {
        return nil, err
    }

    return asn1Bytes, nil
}

func main() {
    // Create the command
    cmd := NewCommand()

    // Execute the command
    if err := cmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
