package Polaris

type Action struct{
	nrl					string
	account				string
	mode				uint
	time				uint			// share action only
	vote				int8
}

const ACTION_MODE_LIKE			= 1
const ACTION_MODE_DISLIKE		= 2
const ACTION_MODE_RESHARE		= 4
const ACTION_MODE_BOOKMARK		= 8