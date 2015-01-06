using Go = import "go.capnp";
$Go.package("factom");
$Go.import("github.com/FactomProject/factom");

@0x832bcc6686a26d56;

struct CommitHeader {
	version			@0 : Int8;
	timeNounce		@1 : UInt64;
	payAmount		@2 : UInt64;
}

struct CommitEntryMsg {
	head 			@0 : CommitHeader;
	hashEntry		@1 : Data;
}

struct CommitChainMsg {
	head 			@0 : CommitHeader;
	hashChain		@1 : Data;
	hashChainEntry 	@2 : Data;
	hashEntry 		@3 : Data;
}

struct CommitMsg $Go.doc("commit msg") {
	union {
		entry 	@0 : CommitEntryMsg;
		chain 	@1 : CommitChainMsg;
	}	
}

struct SignedMsg $Go.doc("signed msg") {
	data :union {
		commit  	@0 : CommitMsg;
		other 		@1 : Void;
	}	

	signedData		@2 : Data;
	publicKey		@3 : Data;
}