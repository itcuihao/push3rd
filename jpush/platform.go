package jpush

// SetPlatform 设置发送平台
func (p *PushRecept) SetPlatform(fs ...string) *PushRecept {

	if p.Platform == nil {
		p.Platform = new(Platform)
	}

	lf := len(fs)
	if lf == 0 {
		p.Platform.Value = "all"
		return p
	}

	pfs := make([]string, 0, len(fs))
	for _, f := range fs {
		pf := getPlatform(f)
		if pf == "" {
			continue
		}
		pfs = append(pfs, pf)
	}
	p.Platform.Value = pfs

	return p
}

func getPlatform(f string) string {
	switch f {
	case PlatformA:
		return PlatformA
	case PlatformI:
		return PlatformI
	}
	return ""
}
