package main

import "strings"

func Home(state *State, req *Request) *Response {
	return NewResponse(200)
}

func Echo(state *State, req *Request) *Response {
	return NewResponse(200).addTextBody(strings.Join(req.path[1:], "/"))
}

func UserAgent(state *State, req *Request) *Response {
	return NewResponse(200).addTextBody(req.headers["User-Agent"])
}

func GetFile(state *State, req *Request) *Response {
	fileName := req.path[1]
	buffer := ReadFile(state.fileDir + "/" + fileName)
	if buffer != nil {
		return NewResponse(200).attachFile(buffer)
	} else {
		return NewResponse(404)
	}
}

func PostFile(state *State, req *Request) *Response {
	fileName := req.path[1]
	WriteFile(state.fileDir+"/"+fileName, req.body)
	return NewResponse(201)
}
