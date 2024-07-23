// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

var UpdateFormTagId = "update"

func UpdateFormInit(formTagId string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_UpdateFormInit_f3cc`,
		Function: `function __templ_UpdateFormInit_f3cc(formTagId){$(` + "`" + `#${formTagId}` + "`" + `)
        .form({
            on: "blur",
            inline : true,
            preventLeaving: true,
            keyboardShortcuts: false,
            fields: {
                Name: {
                    rules: [
                        {
                            type: "size[2..20]"
                        },
                        {
                            type: "regExp[/^[A-Za-z0-9]+$/]",
                            prompt: "{name} must be alphanumeric only"
                        },
                        {
                            type: "regExp[/^.*[^0-9].*$/]",
                            prompt: "{name} cannot be only numbers"
                        }
                    ]
                },
                Role: {
                    rules: [
                        {
                            type   : "exactCount[1]",
                            prompt : "Select a role"
                        },
                    ]
                }
            }
        })
        .api({
            action: "update user", 
            method: "PATCH",
            onSuccess: function(response, element, xhr) {
                $.toast({
                    class: "success",
                    message: response.message,
                    showProgress: "top",
                });

                _.delay(function() {
                    $(` + "`" + `#${formTagId}` + "`" + `).form("reset");
                }, 1000);
            },
            onFailure: function(response, element, xhr) {
                $.toast({
                    class: "error",
                    message: response.message,
                    showProgress: "top"
                });
            }
        })
    ;
}`,
		Call:       templ.SafeScript(`__templ_UpdateFormInit_f3cc`, formTagId),
		CallInline: templ.SafeScriptInline(`__templ_UpdateFormInit_f3cc`, formTagId),
	}
}

func UpdateForm() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(UpdateFormTagId)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/app/server/component/page/home/form.update.templ`, Line: 64, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"ui inverted form\"><div class=\"required field\"><label>Name</label><div class=\"ui inverted transparent left icon input\"><i class=\"user icon\"></i> <input type=\"text\" placeholder=\"Name\" name=\"Name\"></div></div><div class=\"required field\"><label>Role</label><div class=\"ui inverted search clearable selection dropdown\" id=\"Role\"><input type=\"hidden\" name=\"Role\"> <i class=\"dropdown icon\"></i><div class=\"default text\">Select</div><div class=\"menu\"><div class=\"item\" data-value=\"administrator\"><i class=\"user shield icon\"></i>Administrator</div><div class=\"item\" data-value=\"moderator\"><i class=\"users icon\"></i>Moderator</div><div class=\"item\" data-value=\"contributor\"><i class=\"people carry icon\"></i>Contributor</div></div></div></div><div class=\"ui divider\"></div><div class=\"ui middle aligned center aligned grid\"><div class=\"column\"><button class=\"ui vertical animated fluid primary submit button\"><div class=\"hidden content\"><i class=\"sync icon\"></i></div><div class=\"visible content\">Update</div></button></div></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = UpdateFormInit(UpdateFormTagId).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
