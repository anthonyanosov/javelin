local M = {}
local api = vim.api

function M.show_complexity()
	local file = vim.api.nvim_buf_get_name(0)
	local handle = io.popen("javelin -src " .. file)
	local result = handle:read("*a")
	handle:close()

	local ok, data = pcall(vim.fn.json_decode, result)
	if not ok then
		return
	end

	-- clear old virtual text
	local ns = vim.api.nvim_create_namespace("javelin")
	vim.api.nvim_buf_clear_namespace(0, ns, 0, -1)

	for _, f in ipairs(data) do
		vim.api.nvim_buf_set_extmark(0, ns, f.StartLine, 0, {
			virt_text = { { "Complexity: " .. f.Complexity, "Comment" } },
			virt_text_pos = "eol",
		})
	end
end

return M
