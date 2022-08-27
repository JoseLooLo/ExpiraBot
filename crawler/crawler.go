package crawler

import (
    "errors"
    "log"
	"net/url"
	"strings"
    "net/http"
    "io"
    "github.com/PuerkitoBio/goquery"
    expiraBot "github.com/JoseLooLo/ExpiraBot/database"
)

//Get the user books on the BU website
//Receive as input a BU account
//Return a slice of books
func Crawler(login, password string) ([]expiraBot.Books, error) {
    log.Printf("[Info][Crawler] - Getting books from user %s", login)

    bu_url := "https://pergamum.ufsc.br/pergamum/biblioteca_s/php/login_usu.php"

    cookie, err_cookie := getCookie(bu_url)
    if err_cookie != nil {
        return nil, err_cookie
    }

	post_values := url.Values{
		"login": {login},
		"password": {password},
	}

    request, err_post := http.NewRequest("POST", bu_url, strings.NewReader(post_values.Encode()))
    if err_post != nil {
        return nil, errors.New("[Error][Crawler][Crawler] - " + err_post.Error())
    }

	request.Header.Add("Host", "pergamum.ufsc.br")
    request.Header.Add("Origin", "https://pergamum.ufsc.br")
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    // request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
    // request.Header.Add("Referer", "https://pergamum.ufsc.br/pergamum/biblioteca_s/php/login_usu.php?flag=index.php")
    // request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 OPR/90.0.4480.54")
    request.Header.Add("Cookie", cookie)

    client := &http.Client{}
    res, err_request := client.Do(request)
    if err_request != nil {
        return nil, errors.New("[Error][Crawler][Crawler] - " + err_request.Error())
    }
    defer res.Body.Close()

    books, err_parser := htmlParser(res.Body)
    if err_parser != nil {
        return nil, err_parser
    }

    return books, nil
}

//Parse the html request getting the name and expiration date of the user books
func htmlParser(responseBody io.ReadCloser) ([]expiraBot.Books, error) {
    books := []expiraBot.Books{}
    doc, err := goquery.NewDocumentFromReader(responseBody)
	if err != nil {
        return nil, errors.New("[Error][Crawler][htmlParser] - HTML Parser Error")
	}

    invalid_login := false
    //Lets first find if the login and password are correct
    doc.Find(".tab").Each(func(i int, s *goquery.Selection) {
        //If there is a tab class means the login or password are wrong
        invalid_login = true
    })

    if invalid_login {
        return nil, errors.New("[Error][Crawler][htmlParser] - Invalid login or password")
    }


    doc.Find(".txt_azul_11").Each(func(i int, s *goquery.Selection) {
        if (i % 2 == 0) {
            //There is two txt_azul_11 class in the request.
            //The first one is just a whitespace, so we'll skip it
            return
        }
		title := s.Find(".txt_azul").Text()
        //For some reasom there is a "- Livros" after each book name, so we'll just remove it
        title = strings.Replace(title, "- Livros", "", -1)
        title = strings.TrimSpace(title)
        date := s.Next().Text()
        if (title != "") {
            //Just to make sure we have a title
            log.Printf("[Info][Crawler] - Found book: %s - %s ", title, date)
            temp_book := expiraBot.Books{-1, title, date}
            books = append(books, temp_book)
        }
	})
    
    return books, nil
}

//Send a get request just to get a cookie
//We could use a random number instead, but well...
func getCookie(url string) (string, error) {
    temp_cookie_request, err_get := http.Get(url)
    if err_get != nil {
        return "", errors.New("[Error][Crawler][getCookie] - Request Error")
    }
    return "PHPSESSID=" + temp_cookie_request.Cookies()[0].Value, nil
}