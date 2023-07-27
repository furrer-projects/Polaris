package Polaris

type Content struct{
	nrl						string
	// data model specification
	title 					string
	text					string
	description				string
	keywords				string
	resources				[] string
	createdTime				uint
	modifiedTime			uint
	options					uint
	author					string
	target 					string			// empty: profile feed, channel/group id, account address
}