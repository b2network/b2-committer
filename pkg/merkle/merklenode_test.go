package merkle

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
	mt "github.com/txaty/go-merkletree"
)

// newStateRoot outputs: e5f3471831c881c8f34ea32ab485124a09a0f35f010528eb59c51f82e3805df3
// proof outputs: 0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e
// newStateRoot outputs: 34f02ed77cd04345bf84d0f4c8257f0d10701a1f3a8f8c0affc1af9a6301a896
// proof outputs: 0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e
// newStateRoot outputs: 34f02ed77cd04345bf84d0f4c8257f0d10701a1f3a8f8c0affc1af9a6301a896
// proof outputs: 0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e
// newStateRoot outputs: dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e
// proof outputs: 0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e
// newStateRoot outputs: dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e
// proof outputs: 0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e
// newStateRoot outputs: dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e
// proof outputs: 0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e
// newStateRoot outputs: dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e
// proof outputs: 0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e
// newStateRoot outputs: dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e
// proof outputs: 0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e
// newStateRoot outputs: dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e
// proof outputs: 0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e
// newStateRoot outputs: dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e
// proof outputs: 0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e
// stateRootHash:a021b898b4c842021f7c5d2c25fbaeca94b9a9192ccc295ab519c212fd5c2200
// proofRootHash:1492e66e89e186840231850712161255d203b5bbf48d21242f0b51519b5eb3d4
func TestMerkleNode(t *testing.T) {
	newStateRoots := make([]string, 0)
	newStateRoots = append(newStateRoots, "e5f3471831c881c8f34ea32ab485124a09a0f35f010528eb59c51f82e3805df3")
	newStateRoots = append(newStateRoots, "34f02ed77cd04345bf84d0f4c8257f0d10701a1f3a8f8c0affc1af9a6301a896")
	newStateRoots = append(newStateRoots, "34f02ed77cd04345bf84d0f4c8257f0d10701a1f3a8f8c0affc1af9a6301a896")
	newStateRoots = append(newStateRoots, "dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e")
	newStateRoots = append(newStateRoots, "dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e")
	newStateRoots = append(newStateRoots, "dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e")
	newStateRoots = append(newStateRoots, "dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e")
	newStateRoots = append(newStateRoots, "dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e")
	newStateRoots = append(newStateRoots, "dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e")
	newStateRoots = append(newStateRoots, "dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e")
	newStateRoots = append(newStateRoots, "dd813c06390aa63c284c240ad3746e96db28f9ef5e4c4ea76409a10f839ba55e")

	stateBlocks := GenerateBlocks(newStateRoots)
	stateTree, _ := mt.New(nil, stateBlocks)

	stateTreeProofs := stateTree.Proofs
	for i := 0; i < len(stateTreeProofs); i++ {
		ok, _ := stateTree.Verify(stateBlocks[i], stateTreeProofs[i])
		require.Equal(t, ok, true)
	}

	stateRootHash := stateTree.Root
	println("stateRootHash: ", hex.EncodeToString(stateRootHash))
	for i := 0; i < len(stateBlocks); i++ {
		// if hashFunc is nil, use SHA256 by default
		ok, _ := mt.Verify(stateBlocks[i], stateTreeProofs[i], stateRootHash, nil)
		require.Equal(t, ok, true)
	}
	proofs := make([]string, 0)
	proofs = append(proofs, "0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e")
	proofs = append(proofs, "0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e")
	proofs = append(proofs, "0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e")
	proofs = append(proofs, "0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e")
	proofs = append(proofs, "0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e")
	proofs = append(proofs, "0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e")
	proofs = append(proofs, "0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e")
	proofs = append(proofs, "0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e")
	proofs = append(proofs, "0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e")
	proofs = append(proofs, "0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e")

	proofBlocks := GenerateBlocks(proofs)
	proofTree, _ := mt.New(nil, proofBlocks)

	proofTreeProofs := proofTree.Proofs
	for i := 0; i < len(proofTreeProofs); i++ {
		ok, _ := proofTree.Verify(proofBlocks[i], proofTreeProofs[i])
		require.Equal(t, ok, true)
	}

	proofRootHash := proofTree.Root
	println("proofRootHash: ", hex.EncodeToString(proofRootHash))
	for i := 0; i < len(proofBlocks); i++ {
		// if hashFunc is nil, use SHA256 by default
		ok, _ := mt.Verify(proofBlocks[i], proofTreeProofs[i], proofRootHash, nil)
		require.Equal(t, ok, true)
	}
}
