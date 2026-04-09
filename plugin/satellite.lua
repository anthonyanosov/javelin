local ok, satellite = pcall(require, "satellite")
if not ok then
	return
end

vim.api.nvim_create_user_command("Satellite", function()
	satellite.show_complexity()
end, {
	desc = "Show Go function complexity virtual text",
})

vim.api.nvim_create_user_command("SatelliteClear", function()
	satellite.clear()
end, {
	desc = "Clear Satellite virtual text",
})
