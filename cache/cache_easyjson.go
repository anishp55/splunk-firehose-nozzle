// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package cache

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ = json.RawMessage{}
	_ = jlexer.Lexer{}
	_ = jwriter.Writer{}
)

func easyjsonA591d1bcDecodeGithubComCloudfoundryCommunitySplunkFirehoseNozzleCache(in *jlexer.Lexer, out *App) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Name":
			out.Name = string(in.String())
		case "Guid":
			out.Guid = string(in.String())
		case "SpaceName":
			out.SpaceName = string(in.String())
		case "SpaceGuid":
			out.SpaceGuid = string(in.String())
		case "OrgName":
			out.OrgName = string(in.String())
		case "OrgGuid":
			out.OrgGuid = string(in.String())
		case "Environment":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Environment = make(map[string]interface{})
				} else {
					out.Environment = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 interface{}
					v1 = in.Interface()
					(out.Environment)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "SysEnv":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.SysEnv = make(map[string]interface{})
				} else {
					out.SysEnv = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 interface{}
					v2 = in.Interface()
					(out.SysEnv)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "IgnoredApp":
			out.IgnoredApp = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA591d1bcEncodeGithubComCloudfoundryCommunitySplunkFirehoseNozzleCache(out *jwriter.Writer, in App) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Guid\":")
	out.String(string(in.Guid))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"SpaceName\":")
	out.String(string(in.SpaceName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"SpaceGuid\":")
	out.String(string(in.SpaceGuid))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"OrgName\":")
	out.String(string(in.OrgName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"OrgGuid\":")
	out.String(string(in.OrgGuid))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Environment\":")
	if in.Environment == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v3First := true
		for v3Name, v3Value := range in.Environment {
			if !v3First {
				out.RawByte(',')
			}
			v3First = false
			out.String(string(v3Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v3Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"SysEnv\":")
	if in.SysEnv == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v4First := true
		for v4Name, v4Value := range in.SysEnv {
			if !v4First {
				out.RawByte(',')
			}
			v4First = false
			out.String(string(v4Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v4Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"IgnoredApp\":")
	out.Bool(bool(in.IgnoredApp))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v App) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA591d1bcEncodeGithubComCloudfoundryCommunitySplunkFirehoseNozzleCache(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v App) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA591d1bcEncodeGithubComCloudfoundryCommunitySplunkFirehoseNozzleCache(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *App) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA591d1bcDecodeGithubComCloudfoundryCommunitySplunkFirehoseNozzleCache(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *App) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA591d1bcDecodeGithubComCloudfoundryCommunitySplunkFirehoseNozzleCache(l, v)
}
