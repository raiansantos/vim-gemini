function! s:handler(channel, msg)
  let data = json_decode(a:msg)
  call popup_create(split(data.answer, '\n'), {'moved': 'any', 'border': [1,1,1,1], 'padding': [1,3,1,3], 'maxwidth': '120'})
endfunction

function! s:get_visual_selection()
  let [line_start, column_start] = getpos("'<")[1:2]
  let [line_end, column_end] = getpos("'>")[1:2]
  let lines = getline(line_start, line_end)
  if len(lines) == 0
    return ''
  endif
  let lines[-1] = lines[-1][: column_end - 2]
  let lines[0] = lines[0][column_start - 1:]
  return join(lines, "\n")
endfunction

func! gemini#explain() abort
  let ch = ch_open(g:gemini_server, {'mode': 'raw', 'callback': 's:handler'})
  call ch_sendraw(ch, json_encode({'command': 'explain','filetype': &filetype, 'data': s:get_visual_selection()}))
endfunction

func! gemini#debug() abort
  if has('channel')
    let ch = ch_open(g:gemini_server, {'mode': 'raw', 'callback': 's:handler'})
    call ch_sendraw(ch, json_encode({'command': 'debug', 'filetype': &filetype,'data': s:get_visual_selection()}))
  else
    echom 'Channel not enabled'
  endif
endfunction
