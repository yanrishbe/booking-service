package controller

// func login(w http.ResponseWriter, req *http.Request) {
// 	var user model.User
// 	// read request data
// 	_ = json.NewDecoder(req.Body).Decode(&user)
// 	// todo: check whether password and username are ok
//
// 	token, err := createToken(user.ID)
// 	if err != nil {
// 		// todo: correct check
// 	}
// 	json.NewEncoder(w).Encode(model.JwtToken{Token: token})
// }
//
//
// func CreateAuth(userid uint64, td *model.TokenDetails) error {
// 	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
// 	rt := time.Unix(td.RtExpires, 0)
// 	now := time.Now()
//
// 	errAccess := client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
// 	if errAccess != nil {
// 		return errAccess
// 	}
// 	errRefresh := client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
// 	if errRefresh != nil {
// 		return errRefresh
// 	}
// 	return nil
// }
//
//
// func createToken(userID string) (*model.TokenDetails, error) {
// 	// atClaims := jwt.MapClaims{}
// 	// atClaims["authorized"] = true
// 	// atClaims["user_id"] = userID
// 	// atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
// 	//
// 	// at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	// token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
// 	// if err != nil {
// 	// 	return "", err
// 	// }
// 	// return token, nil
//
// 	td := &model.TokenDetails{}
// 	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
// 	td.AccessUuid = uuid.NewV4().String()
//
// 	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
// 	td.RefreshUuid = uuid.NewV4().String()
//
// 	//Creating Access Token
// 	atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = true
// 	atClaims["access_uuid"] = td.AccessUuid
// 	atClaims["user_id"] = userID
// 	atClaims["exp"] = td.AtExpires
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	var err error
// 	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	//Creating Refresh Token
// 	rtClaims := jwt.MapClaims{}
// 	rtClaims["refresh_uuid"] = td.RefreshUuid
// 	rtClaims["user_id"] = userID
// 	rtClaims["exp"] = td.RtExpires
// 	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
// 	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return td, nil
// }
//
//
// func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
// 	params := req.URL.Query()
// 	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("there was an error")
// 		}
// 		return []byte("secret"), nil
// 	})
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		var user model.User
// 		mapstructure.Decode(claims, &user)
// 		json.NewEncoder(w).Encode(user)
// 	} else {
// 		json.NewEncoder(w).Encode(model.Exception{Message: "invalid authorization token"})
// 	}
// }
//
// func TestEndpoint(w http.ResponseWriter, req *http.Request) {
// 	decoded := context.Get(req, "decoded")
// 	var user model.User
// 	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
// 	json.NewEncoder(w).Encode(user)
// }
