import autoload '../autoload/gemini.vim'

if exists('g:loaded_gemini')
  finish
endif
let g:loaded_gemini = 1

command! -range=% GMExplain call gemini.Explain()
command! -range=% GMDebug call gemini.Debug()
