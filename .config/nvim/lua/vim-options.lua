vim.cmd("set expandtab")
vim.cmd("set tabstop=2")
vim.cmd("set softtabstop=2")
vim.cmd("set shiftwidth=2")
vim.cmd("set number")
vim.g.mapleader = " "

vim.keymap.set("n", "<leader>Q", ":qa!<CR>", {})
vim.keymap.set("n", "<C-s>", ":wa<CR>", {}) 


