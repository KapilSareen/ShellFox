package fetch
import(
	"net/http"
	"io"
	"strings"
)

func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	condensedBody := strings.Join(strings.Fields(string(body)), " ")
	return condensedBody, nil
}

