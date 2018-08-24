package server

import(
	"github.com/ugifractal/go-create-react-app/webpack"
)

type User struct {
	Email string
	FirstName string
	LastName string
}

type ViewData struct {
	CurrentUser User
	assetsMapper webpack.AssetsMapper
}

func NewViewData(buildPath string) (ViewData, error) {
	assetsMapper, err := webpack.NewAssetsMapper(buildPath)
	if err != nil {
		return ViewData{}, err
	}

	return ViewData{
		CurrentUser: User{
			Email: "bill@example.com",
			FirstName: "Bill",
			LastName: "Back",
		},
		assetsMapper: assetsMapper,
	}, nil
}

func (d ViewData) Webpack(file string) string {
	return d.assetsMapper(file)
}

