package Polaris

type Profile struct{
	account				string
	picture				string			// nrl or url
	cover				string			// nrl or url
	createdTime			uint
	modifiedTime		uint
	options				uint
}