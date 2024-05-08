if exists('g:loaded_gemini')
  finish
endif
let g:loaded_gemini = 1

command! -range=% GMExplain call gemini#explain()
command! -range=% GMDebug call gemini#debug()
