local M = {}
local ns = vim.api.nvim_create_namespace("javelin")

function M.show_complexity()
	local file = vim.api.nvim_buf_get_name(0)
	local cmd = "javelin -src " .. vim.fn.shellescape(file) .. " 2>&1"
	local handle = io.popen(cmd)
	if not handle then
		vim.notify("Javelin: failed to run analyzer command", vim.log.levels.ERROR)
		return
	end
	local result = handle:read("*a")
	handle:close()

	local ok, data = pcall(vim.fn.json_decode, result)
	if not ok then
		vim.notify("Javelin: failed to parse JSON output", vim.log.levels.ERROR)
		return
	end

	-- clear old virtual text
	vim.api.nvim_buf_clear_namespace(0, ns, 0, -1)

	for _, f in ipairs(data) do
		vim.api.nvim_buf_set_extmark(0, ns, f.StartLine, 0, {
			virt_text = { { "⚡ " .. f.Complexity, "DiagnosticInfo" } },
			virt_text_pos = "eol",
		})
	end
end

return M
