package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type translationRequest struct {
	OriginalText   string `json:"original_text"`
	TargetLanguage string `json:"target_language"`
}

type translationResponse struct {
	TranslatedText string `json:"translated_text"`
}

func SendTranslateHttp(originalText, targetLang string) string {
	requestData := translationRequest{
		OriginalText:   originalText,
		TargetLanguage: targetLang,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("JSON 인코딩 오류:", err)
		return ""
	}

	resp, err := http.Post("http://localhost:8000/api/translate", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("POST 요청 오류:", err)
		return ""
	}
	defer resp.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("응답 읽기 오류:", err)
		return ""
	}

	var translationResponse translationResponse
	err = json.Unmarshal(body, &translationResponse)
	if err != nil {
		fmt.Println("JSON 디코딩 오류:", err)
		return ""
	}

	return translationResponse.TranslatedText
}
