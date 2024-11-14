package middleware

//func UserMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 예: 헤더에서 JWT 토큰 추출 후 유저 ID 확인
//		tokenString := c.GetHeader("Authorization")
//
//		// (예시용) 토큰을 파싱하여 user ID 추출
//		userID, err := parseTokenAndGetUserID(tokenString)
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//			c.Abort()
//			return
//		}
//
//		// 컨텍스트에 user ID 추가
//		c.Set("userID", userID)
//		c.Next()
//	}
//}
//var jwtSecret = []byte("your_secret_key") // 실제 서비스에서는 강력한 키를 사용하세요.
//
//func parseTokenAndGetUserID(tokenString string) (string, error) {
//	// 토큰 파싱 및 검증
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		// 토큰의 서명 알고리즘이 예상과 일치하는지 확인
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, errors.New("unexpected signing method")
//		}
//		return jwtSecret, nil
//	})
//	if err != nil {
//		return "", err
//	}
//
//	// 토큰의 클레임에서 userID 추출
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		userID, ok := claims["userID"].(string)
//		if !ok {
//			return "", errors.New("userID not found in token claims")
//		}
//		return userID, nil
//	}
//
//	return "", errors.New("invalid token")
//}
