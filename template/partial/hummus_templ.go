// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.709
package partial

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

func hummusStyle() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`position:fixed;`)
	templ_7745c5c3_CSSBuilder.WriteString(`bottom:-4vmin;`)
	templ_7745c5c3_CSSBuilder.WriteString(`left:5vmin;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:25vmin;`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:25vmin;`)
	templ_7745c5c3_CSSBuilder.WriteString(`transform:rotate(10deg);`)
	templ_7745c5c3_CSSID := templ.CSSID(`hummusStyle`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Hummus() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{hummusStyle()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xml:space=\"preserve\" xmlns=\"http://www.w3.org/2000/svg\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `template/partial/hummus.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"0 0 300 300\" preserveAspectRatio=\"xMidYMid meet\"><path style=\"fill:#a9b6bd;stroke-width:.0007723;stroke:#000;stroke-opacity:1;stroke-dasharray:none\" d=\"m39.892 116.189-4.903 23.366v29.548l11.766 12.369 62.749 8.246 43.136 15.117h49.022l-14.705-48.1s-36.277-56.348-39.22-56.348c-2.943 0-62.749-4.81-66.674-4.81-3.916-.004-41.171-6.185-41.171 20.612z\" transform=\"translate(-3.847 -42.019) scale(1.35302)\"></path> <path fill=\"#8cc5e7\" d=\"M171.71 83.662c-4.91-10.04-16.783-1.173-19.788-4.324-15.39-15.83-34.005-19.788-48.59-20.52h-8.061c-14.658.366-33.273 4.617-48.59 20.52-3.005 3.151-14.95-5.716-19.788 4.324C21.25 95.388 38.4 96.561 37.52 101.764c-1.686 9.381-.586 23.306.733 37.01 2.052 23.233-15.317 78.858 20.813 101.211 6.816 4.25 25.212 6.596 41.188 6.963h.147c15.977-.367 32.173-2.712 38.99-6.963 36.203-22.353 18.834-77.978 20.96-101.21 1.245-13.705 2.345-27.63.732-37.01-.88-5.204 16.27-6.45 10.627-18.103z\" style=\"stroke-width:.00014658;stroke:#000;stroke-opacity:1;stroke-dasharray:none\" transform=\"translate(-3.847 -42.019) scale(1.35302)\"></path> <g transform=\"translate(22.921 26.953) scale(.9916)\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\"><path fill=\"#e0dedc\" d=\"M143.2 54.3c-33.4 3.9-28.9 38.7-16 50 24 21 49 0 46.2-21.2-2.5-20.4-19.8-30-30.2-28.8z\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\"></path> <circle fill=\"#111212\" cx=\"145.5\" cy=\"84.3\" r=\"11.4\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\"></circle> <circle fill=\"#fff\" cx=\"142.5\" cy=\"79.4\" r=\"3.6\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\"></circle></g> <g style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\"><path fill=\"#b8937f\" d=\"M108.5 107c-16 2.4-21.7 7-20.5 14.2 2 11.8 39.7 10.5 40.9.6 1-8.5-14.1-15.7-20.4-14.8z\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\" transform=\"translate(22.921 26.953) scale(.9916)\"></path> <path d=\"M98.2 111.8c-2.7 9.8 21.7 8.3 21.1 2-.3-3.7-3.6-8.4-12.3-8.2-3.4.1-7.6 1.6-8.8 6.2z\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\" transform=\"translate(22.921 26.953) scale(.9916)\"></path> <path fill=\"#e0dedc\" d=\"M99 127.7c-.9.4-2.4 10.2 2.2 10.7 3.1.3 11.6 1.3 13.6 0 3.9-2.5 3.5-8.5 1.3-10-3.7-2.4-16.1-1.2-17.1-.7Z\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\" transform=\"translate(22.921 26.953) scale(.9916)\"></path></g> <g transform=\"translate(22.921 26.953) scale(.9916)\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\"><path fill=\"#e0dedc\" d=\"M73.6 54.3c33.4 3.9 28.9 38.7 16 50-24 21-49 0-46.2-21.2 2.6-20.4 19.9-30 30.2-28.8z\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\"></path> <circle fill=\"#111212\" cx=\"71.4\" cy=\"84.3\" r=\"11.4\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\"></circle> <circle fill=\"#fff\" cx=\"74.4\" cy=\"79.4\" r=\"3.6\" style=\"stroke:#000;stroke-opacity:1;stroke-width:.0002;stroke-dasharray:none\"></circle></g> <path fill=\"#b8937f\" d=\"M36.86 187.804c-8.061.074-4.104-17.222.88-13.778 2.418 1.686 2.858 5.57 2.858 8.868 0 1.832-1.466 4.837-3.738 4.91zM146.352 240.938c-4.397-6.523-8.354-1.466-14.73 1.759-3.005 1.539 4.983 7.035 13.924 2.931 2.345-1.099 2.272-2.565.806-4.69zM52.177 241.67c4.398-6.522 8.355-1.465 14.731 1.76 3.005 1.539-4.983 7.035-13.924 2.931-2.346-1.1-2.272-2.565-.807-4.69z\" style=\"stroke-width:.00014658;stroke:#000;stroke-opacity:1;stroke-dasharray:none\" transform=\"translate(-3.847 -42.019) scale(1.35302)\"></path> <path fill=\"#3c89bf\" d=\"M158.005 85.421c-.44 1.539 1.54 1.32 2.272 6.083.293 1.759 6.596-2.565 4.03-5.717-2.124-2.638-5.935-1.905-6.302-.366zM40.525 85.421c.44 1.539-1.54 1.32-2.272 6.083-.293 1.759-6.596-2.565-4.031-5.717 2.052-2.638 5.863-1.905 6.303-.366z\" style=\"stroke-width:.00014658;stroke:#000;stroke-opacity:1;stroke-dasharray:none\" transform=\"translate(-3.847 -42.019) scale(1.35302)\"></path> <path style=\"fill:#cfdde2;stroke-width:.00125274;stroke:#000;stroke-opacity:1;stroke-dasharray:none\" d=\"M57.42-5.318c29.683-17.595 78.158-29.658 125.587 31.087 0 0 29.609 91.344 41.936 133.142 3.94 13.354 3.332 19.305 3.332 19.305S156.374 60.382 152.384 59.055c-3.996-1.328-90.536 9.32-107.849 2.662C29.89 81.685 25.23 126.96 25.23 126.96l-7.322 55.922V125.3l8.857-79.719C28.118 24.648 39.367 5.374 57.42-5.318Z\" transform=\"matrix(.97713 0 0 .83425 25.643 23.02)\"></path> <path style=\"fill:#262626;stroke-width:6.26373\" d=\"M191.89 26.157c0 4.316-2.794 8.47-7.154 9.828-51.162 16.017-105.412 16.017-156.574 0-5.312-1.678-8.312-7.466-6.709-12.703.439-1.359 1.078-2.555 1.96-3.557 1.755-2.117 4.354-3.358 7.072-3.358.921 0 1.798.126 2.719.401 47.892 15.02 98.622 15.02 146.515 0 3.633-1.121 7.429.12 9.79 2.957.877 1.002 1.516 2.198 1.96 3.557.257.959.42 1.917.42 2.875z\" transform=\"matrix(.97713 0 0 .83425 25.643 23.02)\"></path> <path style=\"fill:#2b2b2b;stroke-width:6.26373\" d=\"M189.49 19.725c-1.159 1.522-2.8 2.68-4.754 3.276-51.162 16.016-105.412 16.016-156.574 0-1.96-.595-3.596-1.754-4.755-3.276 1.754-2.117 4.354-3.358 7.072-3.358.921 0 1.798.126 2.719.401 47.892 15.02 98.622 15.02 146.515 0 3.626-1.115 7.422.12 9.777 2.957z\" transform=\"matrix(.97713 0 0 .83425 25.643 23.02)\"></path> <path style=\"fill:#262626;stroke-width:6.26373\" d=\"M192.046 51.952a310.746 310.746 0 0 1-171.194 0c-5.825-1.667-9.101-7.46-7.36-12.69 1.735-5.237 7.473-8.05 12.86-6.509a290.849 290.849 0 0 0 160.194 0c5.387-1.54 11.118 1.272 12.86 6.508 1.741 5.23-1.529 11.024-7.36 12.69z\" transform=\"matrix(.97713 0 0 .83425 25.643 23.02)\"></path> <path style=\"fill:#2b2b2b;stroke-width:6.26373\" d=\"M106.446 57.977A304.66 304.66 0 0 1 22.499 46.19c-1.346-.388-2.455-1.247-3.044-2.368-.326-.62-.639-1.585-.282-2.674.564-1.692 2.243-2.832 4.19-2.832.452 0 .903.063 1.341.188a296.72 296.72 0 0 0 163.49 0c.432-.125.889-.188 1.34-.188 1.942 0 3.627 1.14 4.19 2.832.364 1.09.05 2.054-.281 2.674-.595 1.121-1.704 1.98-3.045 2.368a304.705 304.705 0 0 1-83.952 11.788z\" transform=\"matrix(.97713 0 0 .83425 25.643 23.02)\"></path> <path style=\"fill:#cfd8dc\" d=\"M45.365 29.453h2.87v54.554h-2.87z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#a9b6bd\" d=\"M45.365 29.453h1.615v54.554h-1.615z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#cfd8dc\" d=\"M48.504 28.932h-3.408l1.793-5.095.022.043z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#a9b6bd\" d=\"M47.018 28.932h-1.922l1.793-5.095.022.043z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#ff2c46\" d=\"M49.132 29.235v28.248h11.226V29.235z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#429245\" d=\"m90.91 27.063-1.532 9.414H75.459l2.363-9.414z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#d9dada\" d=\"m89.378 36.477-1.538 9.419H73.09l2.369-9.419z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#262626\" d=\"m87.84 45.896-1.531 9.415H70.722l2.368-9.415z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#43a047\" d=\"M60.358 29.235v9.416h16.185l2.368-9.416z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#dfe5e8\" d=\"M60.358 38.651v9.416h13.818l2.367-9.416z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#2d2d2d\" d=\"M60.358 48.067v9.416h11.45l2.368-9.416z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path style=\"fill:#3c7f3c\" d=\"m77.824 27.063 1.087 2.172-2.38.012z\" transform=\"matrix(1 0 0 1 167.264 119.764)\"></path> <path fill=\"#b8937f\" d=\"M161.67 187.804c8.061.074 4.104-17.222-.88-13.778-2.418 1.686-2.858 5.57-2.858 8.868 0 1.832 1.466 4.837 3.738 4.91z\" style=\"stroke-width:.00014658;stroke:#000;stroke-opacity:1;stroke-dasharray:none\" transform=\"translate(-3.847 -42.019) scale(1.35302)\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}