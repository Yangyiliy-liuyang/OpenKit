package logger

var DEBUG = false

func Error(err error) Field {
	return Field{
		"error",
		err.Error(),
	}
}

// SafeString 非DEBUG模式下，屏蔽敏感信息
func SafeString(key string, val string) Field {
	if DEBUG {
		return Field{
			key,
			val,
		}
	} else {
		return Field{
			key,
			"***",
		}
	}
}

func String(key string, val string) Field {
	return Field{
		key,
		val,
	}
}

func Int64(key string, val int64) Field {
	return Field{
		key,
		val,
	}
}
func Int32(key string, val int32) Field {
	return Field{
		key,
		val,
	}
}

func Int(key string, val int) Field {
	return Field{
		key,
		val,
	}
}
