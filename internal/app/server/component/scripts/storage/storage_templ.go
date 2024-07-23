// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package storage

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Init() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_Init_f92a`,
		Function: `function __templ_Init_f92a(){const Storage = {
        MasterKey: "dsgo",
        Key: {},
        Init() {
            let storage = localStorage.getItem(this.MasterKey);

            if (storage == null) {
                localStorage.setItem(this.MasterKey, JSON.stringify({}));
            }
        },
        Put(key, value) {
            let storage = localStorage.getItem(this.MasterKey);

            storage = JSON.parse(storage)

            storage[key] = value;

            localStorage.setItem(this.MasterKey, JSON.stringify(storage));
        },
        Get(key) {
            let storage = localStorage.getItem(this.MasterKey);

            storage = JSON.parse(storage)

            return _.get(storage, key, null);
        },
        Delete(key) {
            let storage = localStorage.getItem(this.MasterKey);

            storage = JSON.parse(storage)

            delete storage[key];

            localStorage.setItem(this.MasterKey, JSON.stringify(storage));
        },
        async ClearSession() {
            localStorage.removeItem(this.MasterKey);
            cookieStore.delete(this.MasterKey);
        },
        async Clear() {
            localStorage.clear();

            let cookies = await cookieStore.getAll();

            _.each(cookies, function(cookie) {
                cookieStore.delete(cookie);
            });
        }
    }

    Storage.Init();

    window.Storage = Storage
}`,
		Call:       templ.SafeScript(`__templ_Init_f92a`),
		CallInline: templ.SafeScriptInline(`__templ_Init_f92a`),
	}
}

func Storage() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		return templ_7745c5c3_Err
	})
}
