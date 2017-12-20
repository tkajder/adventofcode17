package knothash

import "encoding/hex"

type knotHash struct {
	data     [256]byte
	currPos  byte
	skipSize uint
}

func new() *knotHash {
	kh := &knotHash{}
	kh.reset()
	return kh
}

func (kh *knotHash) reset() {
	kh.data = [256]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
		61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
		81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
		91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
		101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
		111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
		121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
		131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
		141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
		151, 152, 153, 154, 155, 156, 157, 158, 159, 160,
		161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
		171, 172, 173, 174, 175, 176, 177, 178, 179, 180,
		181, 182, 183, 184, 185, 186, 187, 188, 189, 190,
		191, 192, 193, 194, 195, 196, 197, 198, 199, 200,
		201, 202, 203, 204, 205, 206, 207, 208, 209, 210,
		211, 212, 213, 214, 215, 216, 217, 218, 219, 220,
		221, 222, 223, 224, 225, 226, 227, 228, 229, 230,
		231, 232, 233, 234, 235, 236, 237, 238, 239, 240,
		241, 242, 243, 244, 245, 246, 247, 248, 249, 250,
		251, 252, 253, 254, 255}

	kh.currPos = 0
	kh.skipSize = 0
}

func (kh *knotHash) knot(length byte) {
	// Find start and end of length with overflow
	lengthStart := kh.currPos
	lengthEnd := lengthStart + length

	// Reverse length with wrap around indexing
	for i := byte(0); i < length/2; i++ {
		// Calculate swap indices with overflowed length
		swapStart := lengthStart + i
		swapEnd := lengthEnd - i - 1

		// Perform swap
		temp := kh.data[swapStart]
		kh.data[swapStart] = kh.data[swapEnd]
		kh.data[swapEnd] = temp
	}

	// Move current position forward by length and skipSize with overflow
	kh.currPos = kh.currPos + length + byte(kh.skipSize)

	// Increment skipSize
	kh.skipSize++
}

func (kh *knotHash) Hash(p []byte) string {
	suffix := getBytesSuffix()

	// Compute 64 rounds of knotting on input and suffix
	for round := 0; round < 64; round++ {
		for _, b := range p {
			kh.knot(b)
		}

		for _, b := range suffix {
			kh.knot(b)
		}
	}

	// Calculate dense hashes from sparse hash
	denseHashes := kh.calculateDenseHashes()

	// reset after hash
	kh.reset()

	return hex.EncodeToString(denseHashes)
}

func (kh *knotHash) calculateDenseHashes() []byte {
	denseHashes := make([]byte, 16)

	for block := 0; block < 16; block++ {
		denseHash := byte(0)
		offset := block * 16

		// Calculate dense hash for block
		for blockMember := 0; blockMember < 16; blockMember++ {
			denseHash ^= kh.data[offset+blockMember]
		}

		denseHashes[block] = denseHash
	}

	return denseHashes
}

func getBytesSuffix() []byte {
	return []byte{17, 31, 73, 47, 23}
}
