package types

import (
	"encoding/json"
	"testing"
)

func TestUnmarshelWalData(t *testing.T) {
	t.Helper()
	dat := []byte(`	{
		"nextlsn": "2/7EEE60A8",
		"change": [{
			"kind": "update",
			"schema": "public",
			"table": "user_profile",
			"columnnames": ["user_id", "title", "full_name", "about_me", "nationality", "date_of_birth", "gender", "resident_card_no", "phone_mobile", "phone_home", "subject", "email", "last_login_date", "is_searchable", "is_active", "is_locked", "bkg_img", "photo", "country_id", "city_id", "state_id", "created_date", "changed_date", "deleted_date", "groups", "display_name", "disabilities", "marital_status", "driving_license_no", "driving_license_date", "driving_license_type", "passport_no", "tax_no", "vat_no", "street_address", "postal_code", "phone_work", "metadata", "is_contact_info_visible", "primary_language", "device_token", "access_token_count", "picture_url", "dependents_count"],
			"columntypes": ["bigint", "character varying(128)", "character varying(512)", "character varying(8192)", "character varying(2)", "timestamp without time zone", "character varying(10)", "character varying(32)", "character varying(64)", "character varying(64)", "character varying(512)", "character varying(256)", "timestamp without time zone", "boolean", "boolean", "boolean", "character varying(1024)", "bigint", "bigint", "bigint", "bigint", "timestamp without time zone", "timestamp without time zone", "timestamp without time zone", "character varying(256)", "character varying(256)", "character varying(256)", "character varying(256)", "character varying(32)", "timestamp without time zone", "character varying(32)", "character varying(32)", "character varying(32)", "character varying(32)", "character varying(512)", "character varying(10)", "character varying(64)", "character varying(8192)", "boolean", "bigint", "character varying(512)", "bigint", "character varying(256)", "integer"],
			"columnvalues": [0, null, "hamdi", null, null, null, null, null, null, null, null, "", null, null, true, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, false, null, null, null, null, null],
			"oldkeys": {
				"keynames": ["user_id"],
				"keytypes": ["bigint"],
				"keyvalues": [0]
			}
		}]
	}`)

	var wData WalData
	if err := json.Unmarshal(dat, &wData); err != nil {
		t.Fail()
	}
	// v := WalData{
	// 	Nextlsn: "2/7EEE60A8",
	// 	Change: []Change{
	// 		Change{
	// 			Kind:         "update",
	// 			Schema:       "public",
	// 			Table:        "user_profile",
	// 			Columnnames:  []string{"user_id", "title", "full_name", "about_me", "nationality", "date_of_birth", "gender", "resident_card_no", "phone_mobile", "phone_home", "subject", "email", "last_login_date", "is_searchable", "is_active", "is_locked", "bkg_img", "photo", "country_id", "city_id", "state_id", "created_date", "changed_date", "deleted_date", "groups", "display_name", "disabilities", "marital_status", "driving_license_no", "driving_license_date", "driving_license_type", "passport_no", "tax_no", "vat_no", "street_address", "postal_code", "phone_work", "metadata", "is_contact_info_visible", "primary_language", "device_token", "access_token_count", "picture_url", "dependents_count"},
	// 			Columntypes:  []string{"bigint", "character varying(128)", "character varying(512)", "character varying(8192)", "character varying(2)", "timestamp without time zone", "character varying(10)", "character varying(32)", "character varying(64)", "character varying(64)", "character varying(512)", "character varying(256)", "timestamp without time zone", "boolean", "boolean", "boolean", "character varying(1024)", "bigint", "bigint", "bigint", "bigint", "timestamp without time zone", "timestamp without time zone", "timestamp without time zone", "character varying(256)", "character varying(256)", "character varying(256)", "character varying(256)", "character varying(32)", "timestamp without time zone", "character varying(32)", "character varying(32)", "character varying(32)", "character varying(32)", "character varying(512)", "character varying(10)", "character varying(64)", "character varying(8192)", "boolean", "bigint", "character varying(512)", "bigint", "character varying(256)", "integer"},
	// 			Columnvalues: []interface{}{0, nil, "hamdi", nil, nil, nil, nil, nil, nil, nil, nil, "", nil, nil, true, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, false, nil, nil, nil, nil, nil},
	// 			Oldkeys: OldKeys{
	// 				Keynames:  []string{"user_id"},
	// 				Keytypes:  []string{"bigint"},
	// 				Keyvalues: []int64{0},
	// 			},
	// 		},
	// 	},
	// }

	// assert.Equal(t, wData, v)

}
