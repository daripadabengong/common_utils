package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type EntityID struct {
	value uuid.UUID
}

type PhoneNumber struct {
	value string
}

type EmailAddress struct {
	value string
}

type NullableString struct {
	value string
}

type RequiredString struct {
	value string
}

func (v EntityID) GetValue() uuid.UUID    { return v.value }
func (v PhoneNumber) GetValue() string    { return v.value }
func (v EmailAddress) GetValue() string   { return v.value }
func (v NullableString) GetValue() string { return v.value }
func (v RequiredString) GetValue() string { return v.value }

func NewEntityID(value uuid.UUID) (EntityID, error) {
	if value == uuid.Nil {
		value = uuid.New()
	}
	return EntityID{value: value}, nil
}

func NewPhoneNumber(value string) (PhoneNumber, error) {
	if value == "" {
		return PhoneNumber{}, errors.New("phone number can't be empty")
	}

	// Check if the phone number starts with "+"
	if !strings.HasPrefix(value, "+") {
		return PhoneNumber{}, errors.New("phone number must start with '+'")
	}

	// Remove the "+" for easier processing
	phoneWithoutPlus := value[1:]

	// Match and extract the country code
	regex := regexp.MustCompile(`^\d{1,3}`) // Matches up to 3 digits at the start
	matches := regex.FindStringSubmatch(phoneWithoutPlus)
	if len(matches) == 0 {
		return PhoneNumber{}, errors.New("invalid phone number format")
	}

	countryCode := matches[0]

	fmt.Println(countryCode)

	// Validate the country code against the map
	if _, exists := countryCodeMap[countryCode]; !exists {
		return PhoneNumber{}, errors.New("invalid or unsupported country code")
	}
	return PhoneNumber{value: value}, nil
}

func NewEmailAddress(value string) (EmailAddress, error) {
	if value == "" {
		return EmailAddress{}, errors.New("email address can't be empty")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(value) {
		return EmailAddress{}, errors.New("invalid email address")
	}
	return EmailAddress{value: value}, nil
}

func NewNullableString(value string) (NullableString, error) {
	return NullableString{value: value}, nil
}

func NewRequiredString(fieldName, value string) (RequiredString, error) {
	if value == "" {
		return RequiredString{}, fmt.Errorf("%s can't be empty", fieldName)
	}
	return RequiredString{value: value}, nil
}

var countryCodeMap = map[string]string{
	"1":   "US", // United States/Canada
	"7":   "RU", // Russia/Kazakhstan
	"20":  "EG", // Egypt
	"27":  "ZA", // South Africa
	"30":  "GR", // Greece
	"31":  "NL", // Netherlands
	"32":  "BE", // Belgium
	"33":  "FR", // France
	"34":  "ES", // Spain
	"36":  "HU", // Hungary
	"39":  "IT", // Italy
	"40":  "RO", // Romania
	"41":  "CH", // Switzerland
	"43":  "AT", // Austria
	"44":  "GB", // United Kingdom
	"45":  "DK", // Denmark
	"46":  "SE", // Sweden
	"47":  "NO", // Norway
	"48":  "PL", // Poland
	"49":  "DE", // Germany
	"51":  "PE", // Peru
	"52":  "MX", // Mexico
	"53":  "CU", // Cuba
	"54":  "AR", // Argentina
	"55":  "BR", // Brazil
	"56":  "CL", // Chile
	"57":  "CO", // Colombia
	"58":  "VE", // Venezuela
	"60":  "MY", // Malaysia
	"61":  "AU", // Australia
	"62":  "ID", // Indonesia
	"63":  "PH", // Philippines
	"64":  "NZ", // New Zealand
	"65":  "SG", // Singapore
	"66":  "TH", // Thailand
	"81":  "JP", // Japan
	"82":  "KR", // South Korea
	"84":  "VN", // Vietnam
	"86":  "CN", // China
	"90":  "TR", // Turkey
	"91":  "IN", // India
	"92":  "PK", // Pakistan
	"93":  "AF", // Afghanistan
	"94":  "LK", // Sri Lanka
	"95":  "MM", // Myanmar
	"98":  "IR", // Iran
	"211": "SS", // South Sudan
	"212": "MA", // Morocco
	"213": "DZ", // Algeria
	"216": "TN", // Tunisia
	"218": "LY", // Libya
	"220": "GM", // Gambia
	"221": "SN", // Senegal
	"222": "MR", // Mauritania
	"223": "ML", // Mali
	"224": "GN", // Guinea
	"225": "CI", // Ivory Coast
	"226": "BF", // Burkina Faso
	"227": "NE", // Niger
	"228": "TG", // Togo
	"229": "BJ", // Benin
	"230": "MU", // Mauritius
	"231": "LR", // Liberia
	"232": "SL", // Sierra Leone
	"233": "GH", // Ghana
	"234": "NG", // Nigeria
	"235": "TD", // Chad
	"236": "CF", // Central African Republic
	"237": "CM", // Cameroon
	"238": "CV", // Cape Verde
	"239": "ST", // Sao Tome and Principe
	"240": "GQ", // Equatorial Guinea
	"241": "GA", // Gabon
	"242": "CG", // Congo
	"243": "CD", // Democratic Republic of the Congo
	"244": "AO", // Angola
	"245": "GW", // Guinea-Bissau
	"246": "IO", // British Indian Ocean Territory
	"248": "SC", // Seychelles
	"249": "SD", // Sudan
	"250": "RW", // Rwanda
	"251": "ET", // Ethiopia
	"252": "SO", // Somalia
	"253": "DJ", // Djibouti
	"254": "KE", // Kenya
	"255": "TZ", // Tanzania
	"256": "UG", // Uganda
	"257": "BI", // Burundi
	"258": "MZ", // Mozambique
	"260": "ZM", // Zambia
	"261": "MG", // Madagascar
	"262": "RE", // Réunion
	"263": "ZW", // Zimbabwe
	"264": "NA", // Namibia
	"265": "MW", // Malawi
	"266": "LS", // Lesotho
	"267": "BW", // Botswana
	"268": "SZ", // Eswatini
	"269": "KM", // Comoros
	"290": "SH", // Saint Helena
	"291": "ER", // Eritrea
	"297": "AW", // Aruba
	"298": "FO", // Faroe Islands
	"299": "GL", // Greenland
	"350": "GI", // Gibraltar
	"351": "PT", // Portugal
	"352": "LU", // Luxembourg
	"353": "IE", // Ireland
	"354": "IS", // Iceland
	"355": "AL", // Albania
	"356": "MT", // Malta
	"357": "CY", // Cyprus
	"358": "FI", // Finland
	"359": "BG", // Bulgaria
	"370": "LT", // Lithuania
	"371": "LV", // Latvia
	"372": "EE", // Estonia
	"373": "MD", // Moldova
	"374": "AM", // Armenia
	"375": "BY", // Belarus
	"376": "AD", // Andorra
	"377": "MC", // Monaco
	"378": "SM", // San Marino
	"380": "UA", // Ukraine
	"381": "RS", // Serbia
	"382": "ME", // Montenegro
	"383": "XK", // Kosovo
	"385": "HR", // Croatia
	"386": "SI", // Slovenia
	"387": "BA", // Bosnia and Herzegovina
	"389": "MK", // North Macedonia
	"420": "CZ", // Czech Republic
	"421": "SK", // Slovakia
	"423": "LI", // Liechtenstein
	"500": "FK", // Falkland Islands
	"501": "BZ", // Belize
	"502": "GT", // Guatemala
	"503": "SV", // El Salvador
	"504": "HN", // Honduras
	"505": "NI", // Nicaragua
	"506": "CR", // Costa Rica
	"507": "PA", // Panama
	"508": "PM", // Saint Pierre and Miquelon
	"509": "HT", // Haiti
	"590": "GP", // Guadeloupe
	"591": "BO", // Bolivia
	"592": "GY", // Guyana
	"593": "EC", // Ecuador
	"594": "GF", // French Guiana
	"595": "PY", // Paraguay
	"596": "MQ", // Martinique
	"597": "SR", // Suriname
	"598": "UY", // Uruguay
	"599": "AN", // Netherlands Antilles
}
