---@module 'lazy'
---@type LazySpec
return {
	{
		"zbirenbaum/copilot.lua",
		---@module 'copilot'
		---@type CopilotConfig
		---@diagnostic disable-next-line: missing-fields
		opts = {
			---@diagnostic disable-next-line: missing-fields
			suggestion = {
				enabled = false,
			},
		},
		optional = true,
	},
}
