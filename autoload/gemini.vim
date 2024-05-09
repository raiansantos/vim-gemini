vim9script

var ch: channel

def Handler(channel: any, msg: string)
  var data = json_decode(msg)
  call popup_create(split(data.answer, '\n'), {'moved': 'any', 'border': [1, 1, 1, 1], 'padding': [1, 3, 1, 3], 'maxwidth': 120})
enddef

def GetVisualSelection(): string
  var [line_start, column_start] = getpos("'<")[1 : 2]
  var [line_end, column_end] = getpos("'>")[1 : 2]
  var lines = getline(line_start, line_end)
  if len(lines) == 0
    return ''
  endif
  lines[-1] = lines[-1][ : column_end - 2]
  lines[0] = lines[0][column_start - 1 : ]
  return join(lines, "\n")
enddef

export def Explain()
  if ch_status(ch) != 'open'
    ch = ch_open(g:gemini_server, {'mode': 'raw', 'callback': Handler, 'waittime': 1000})
  endif

  if ch_status(ch) != 'open'
    echom "Channel is not open"
    echom ch_info(ch)
    return
  endif

  ch_sendraw(ch, json_encode({'command': 'explain', 'filetype': &filetype, 'data': GetVisualSelection()}))
enddef

export def Debug()
  if ch_status(ch) != 'open'
    ch = ch_open(g:gemini_server, {'mode': 'raw', 'callback': Handler, 'waittime': 1000})
  endif

  if ch_status(ch) != 'open'
    echom "Channel is not open"
    echom ch_info(ch)
    return
  endif

  ch_sendraw(ch, json_encode({'command': 'debug', 'filetype': &filetype, 'data': GetVisualSelection()}))
enddef
