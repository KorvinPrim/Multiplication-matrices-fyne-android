package main

/*
This training program focuses on implementing resizing and positioning
elements using the Go language of the Fyne library. You will learn how
to create a user-friendly registration form, where elements such as
text fields, buttons, labels, and checkboxes can be adjusted in size
and positioned precisely within the form's layout. This program will
provide hands-on exercises and guidance on leveraging the Fyne library's
capabilities to create a dynamic and visually appealing user interface
for your registration form.
*/

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
	"strings"
)

var number_items float32 = 1
var todos = binding.NewUntypedList()

func collectres(result [4][4]float64, course_decision [][4][4]float64, diu []interface{}) {
	res_for_w := ""
	todos.Append(NewTodo(""))
	todos.Append(NewTodo("Decision"))
	fmt.Println("The course decision of multiplication of matrix")
	qvestion := map[int]string{}
	for i := 0; i < len(diu); i++ {
		str_d := strings.Split(fmt.Sprintf("%v", diu[i]), "  - ")[0]
		parse_d := fmt.Sprintf("%v", str_d)
		qvestion[i] = parse_d
		fmt.Println(qvestion)
	}

	todos.Append(NewTodo(fmt.Sprintf("%v\n", qvestion[0])))

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			res_for_w += fmt.Sprintf("%0.3f\t", course_decision[0][i][j])
		}
		res_for_w += "\n"
	}
	todos.Append(NewTodo(res_for_w))
	todos.Append(NewTodo(""))

	res_for_w = ""

	for m := 1; m < len(course_decision); m++ {
		if m%2 != 0 && m != 0 {
			if m == 1 {
				todos.Append(NewTodo(qvestion[m] + "\n"))
			} else {
				todos.Append(NewTodo(qvestion[m/2] + "\n"))
			}
		} else {
			todos.Append(NewTodo("Result of multiplication"))
		}
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				res_for_w += fmt.Sprintf("%0.3f\t", course_decision[m][i][j])
			}
			res_for_w += "\n"
		}
		fmt.Println(course_decision)
		todos.Append(NewTodo(res_for_w))
		todos.Append(NewTodo(""))
		res_for_w = ""

	}
	todos.Append(NewTodo("The final results of matrix multiplication"))
	res_for_w = ""

	fmt.Println(result)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			res_for_w += fmt.Sprintf("%0.3f\t", result[i][j])
		}
		res_for_w = fmt.Sprintf("%v\n", res_for_w)
	}
	res_for_w = fmt.Sprintf(res_for_w)
	todos.Append(NewTodo(res_for_w))
	todos.Append(NewTodo(""))
	todos.Append(NewTodo(""))
}

func main() {
	a := app.New()
	w := a.NewWindow("MHRT matrices")
	w.Resize(fyne.NewSize(320, 480))
	a.Settings().SetTheme(theme.DarkTheme())

	data := []Todo{}

	for _, t := range data {
		todos.Append(t)
	}

	label := canvas.NewText("Добавьте матрицу", color.White)
	label.TextSize = 20
	label.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	label.Resize(fyne.NewSize(100, 30))
	label.Move(fyne.NewPos(60, 5))

	newtodoDescTxt := widget.NewEntry()
	newtodoDescTxt.PlaceHolder = "Формат ввода: q|h|b x|y|z 1.57"
	submit := widget.NewButton("Добавить", func() {
		list_data := strings.Split(fmt.Sprintf("%v", newtodoDescTxt.Text), " ")

		if len(list_data) == 3 {
			if list_data[0] == "q" || list_data[0] == "h" || list_data[0] == "b" {
				if list_data[1] == "x" || list_data[1] == "y" || list_data[1] == "z" {

					if _, err := strconv.ParseFloat(list_data[2], 64); err == nil {
						todos.Append(NewTodo(newtodoDescTxt.Text))
						newtodoDescTxt.Text = ""
					}

				}
			}
		}
	})

	submit.Resize(fyne.NewSize(290, 50))
	submit.Move(fyne.NewPos(10, number_items*60))

	multiplications := widget.NewButton("Рассчитать умножение", func() {
		matrix_for_multiplication := [][4][4]float64{}
		diu, _ := todos.Get()

		for i := 0; i < len(diu); i++ {
			str_d := strings.Split(fmt.Sprintf("%v", diu[i]), "  - ")[0]
			parse_d := strings.Split(fmt.Sprintf("%v", str_d), " ")
			if parse_d[0] == "q" {
				switch parse_d[1] {
				case "x":
					f, _ := strconv.ParseFloat(parse_d[2], 64)
					matrix_for_multiplication = append(matrix_for_multiplication, rotX(f))
				case "y":
					f, _ := strconv.ParseFloat(parse_d[2], 64)
					matrix_for_multiplication = append(matrix_for_multiplication, rotY(f))
				case "z":
					f, _ := strconv.ParseFloat(parse_d[2], 64)
					matrix_for_multiplication = append(matrix_for_multiplication, rotZ(f))
				}
			}
			if parse_d[0] == "h" || parse_d[0] == "b" {
				switch parse_d[1] {
				case "x":
					f, _ := strconv.ParseFloat(parse_d[2], 64)
					matrix_for_multiplication = append(matrix_for_multiplication, smX(f))
				case "y":
					f, _ := strconv.ParseFloat(parse_d[2], 64)
					matrix_for_multiplication = append(matrix_for_multiplication, smY(f))
				case "z":
					f, _ := strconv.ParseFloat(parse_d[2], 64)
					matrix_for_multiplication = append(matrix_for_multiplication, smZ(f))
				}

			}

		}

		result, course_decision := MultiplyMultipleMatrices(matrix_for_multiplication)

		collectres(result, course_decision, diu)

	})
	multiplications.Resize(fyne.NewSize(290, 50))
	multiplications.Move(fyne.NewPos(10, number_items*60))

	w.SetContent(
		container.NewBorder(
			container.NewBorder(
				nil, // TOP
				nil, // BOTTOM
				nil, // Left
				// RIGHT ↓
				submit,
				newtodoDescTxt,
				// take the rest of the space
			),

			// TOP of the container
			container.NewBorder(
				nil, // TOP
				nil, // BOTTOM
				nil, // Left
				// RIGHT ↓

				multiplications,
				// take the rest of the space
			),
			nil, // Right
			nil, // Left
			// the rest will take all the rest of the space
			widget.NewListWithData(
				// the binding.List type
				todos,
				// func that returns the component structure of the List Item
				// exactly the same as the Simple List
				func() fyne.CanvasObject {
					return container.NewBorder(
						nil, nil, nil,
						// left of the border
						//widget.NewCheck("", func(b bool) {}),
						// takes the rest of the space
						widget.NewLabel(""),
					)
				},
				// func that is called for each item in the list and allows
				// but this time we get the actual DataItem we need to cast
				func(di binding.DataItem, o fyne.CanvasObject) {
					ctr, _ := o.(*fyne.Container)
					// ideally we should check `ok` for each one of those casting
					// but we know that they are those types for sure
					l := ctr.Objects[0].(*widget.Label)
					l.Resize(fyne.NewSize(290, 50))

					diu, _ := di.(binding.Untyped).Get()
					todo := diu.(Todo)
					//todo := NewTodoFromDataItem(di)
					l.SetText(todo.Description)

				}),
		),
	)
	w.ShowAndRun()

}

//fyne package -os windows -icon myapp.png
//fyne package -os android -appID com.example.myapp -icon icon.png ANDROID_NDK_HOME=C:\Users\Korvin\go\src\fyne.io\fyne\v2\android-ndk-r25c
//fyne package -os android -appID com.example.myapp -icon icon.png ANDROID_NDK_HOME=C:\Users\Korvin\Library\Android\sdk
