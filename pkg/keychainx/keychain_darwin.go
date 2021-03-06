// +build darwin

/*
Copyright 2019 Adobe
All Rights Reserved.

NOTICE: Adobe permits you to use, modify, and distribute this file in
accordance with the terms of the Adobe license agreement accompanying
it. If you have received this file from a source other than Adobe,
then your use, modification, or distribution of it requires the prior
written permission of Adobe.
*/

package keychainx

import (
	"github.com/keybase/go-keychain"
)

func Save(label, user, password string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassInternetPassword)
	item.SetLabel(label)
	item.SetAccount(user)
	item.SetData([]byte(password))

	return keychain.AddItem(item)
}

// Load credentials with a given label
func Load(label string) (string, string, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassInternetPassword)
	query.SetLabel(label)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnAttributes(true)
	query.SetReturnData(true)

	results, err := keychain.QueryItem(query)
	if err != nil {
		return "", "", err
	}

	for _, r := range results {
		return string(r.Account), string(r.Data), nil
	}

	return "", "", ErrNotFound
}
