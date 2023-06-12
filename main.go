package main

import (
	"fmt"
	"strconv"
)

func crc16(data []byte) uint16 {
	// Polinomial pembagi CRC-16 (0x8005) dalam bentuk polinomial 17-bit
	polynomial := uint16(0x8005)
	var crc uint16

	for _, b := range data {
		crc ^= uint16(b) << 8
		for i := 0; i < 8; i++ {
			if crc&(1<<15) != 0 {
				crc = (crc << 1) ^ polynomial
			} else {
				crc <<= 1
			}
		}
	}

	return crc
}

func verifyCRC(str string, crc uint16) bool {
	data := []byte(str)
	calculatedCRC := crc16(data)

	return calculatedCRC == crc
}

func main() {
	// Contoh string
	str := "203894023702347293874872394723987482379423849832"

	// Mengubah string menjadi byte array
	data := []byte(str)

	// Menghitung CRC-16
	crc := crc16(data)

	// Menampilkan nilai CRC
	fmt.Printf("CRC-16: %04x\n", crc)

	expectedCRCStr := "4373" // Nilai CRC yang diharapkan dalam bentuk string heksadesimal

	// Mengonversi nilai CRC yang diharapkan menjadi uint16
	expectedCRC, err := strconv.ParseUint(expectedCRCStr, 16, 16)
	if err != nil {
		fmt.Println("Invalid expected CRC value:", err)
		return
	}

	// Memverifikasi CRC-16
	isValid := verifyCRC(str, uint16(expectedCRC))

	// Menampilkan hasil verifikasi
	if isValid {
		fmt.Println("CRC-16 valid. String belum mengalami perubahan.")
	} else {
		fmt.Println("CRC-16 tidak valid. String telah mengalami perubahan atau CRC salah.")
	}
}
