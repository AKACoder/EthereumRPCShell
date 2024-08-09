package enum

type addressable interface {
    Get(elName string) *Element
    setEnum(string, int, interface{})
}

type AddressableEnum struct {
    inited bool

    nameMapping  map[string]interface{}
    valueMapping map[int]interface{}
}

func (a *AddressableEnum) setEnum(elName string, val int, el interface{}) {
    if !a.inited {
        a.nameMapping = make(map[string]interface{})
        a.valueMapping = make(map[int]interface{})

        a.inited = true
    }

    a.nameMapping[elName] = el
    a.valueMapping[val] = el
}

func (a *AddressableEnum) GetByVal(val int) *Element {
    el, exist := a.valueMapping[val]
    if !exist {
        return nil
    }

    ret, ok := el.(Element)
    if ok {
        return &ret
    } else {
        return nil
    }
}

func (a *AddressableEnum) Get(elName string) *Element {
    el, exist := a.nameMapping[elName]
    if !exist {
        return nil
    }

    ret, ok := el.(Element)
    if ok {
        return &ret
    } else {
        return nil
    }
}
