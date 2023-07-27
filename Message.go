package Polaris

type Message struct{
	nrl						string
	// data model specification
	text					string
	resources				[] string
	createdTime				uint
	modifiedTime			uint
	options					uint
	author					string
	target 					string			// empty: profile feed, channel/group id, account address
}