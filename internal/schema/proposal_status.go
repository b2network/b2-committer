package schema

const (
	ProposalVotingStatus  = 0 // 0 Proposer can submit proof and vote for the proposal
	ProposalPendingStatus = 1 // 1 Proposer can submit decentralized storage tx for the proposal
	ProposalCommitting    = 2 // 2 Proposer can submit bitcoin tx for the proposal
	ProposalSucceedStatus = 3 // 3 Proposal is success
	ProposalTimeoutStatus = 4 // 4 Proposal is timed out
)
