package phonegen

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
Telefonia fixa
Para as operadoras de telefonia fixa é destinado a faixa de prefixos que têm números entre 2 a 5 para o primeiro dígito (N8).

Telefonia móvel
Para as operadoras de telefonia móvel é destinado a faixa entre 6 a 9 (em SP utiliza-se o 5 também) para o primeiro dígito (N8). Desta faixa é possível identificar a peculiaridade de cada um desses dígitos (6-9):

- 9 utilizado para as Bandas A (96 a 99) e B (91 a 94);

- 8 utilizado para as bandas D e E;

- 7 destinado a celulares e Trunking (Nextel);

- 6 utilizado para as bandas A, B, D e E* (utilizado em SP para telefonia móvel a partir de 2008).

Não Geográficos
Códigos não geográficos são códigos utilizáveis em todo o território do país. Eles foram definidos pela Anatel, são eles:

900: onde o originador é responsável pelo pagamento da chamada;
800: onde o destino é quem paga pela chamada;
500: destinado ao registro de intenção de doação (valor máximo de R$ 30,00);
300: originador se responsabiliza pelo pagamento da chamada.
Há ainda os códigos de acesso a serviços de utilidade pública, que são formados por três dígitos. Todos têm como primeiro dígito o número 1, caracterizadas por serem chamadas gratuitas (quando originadas de um telefone fixo) para a maioria dos serviços.
*/

var (
	areaCodes = map[string]map[string]string{
		// Sudeste
		"11": {"São Paulo": "Região Metropolitana de São Paulo"},
		"12": {"São Paulo": "São José dos Campos e Região"},
		"13": {"São Paulo": "Região Metropolitana da Baixada Santista"},
		"14": {"São Paulo": "Bauru, Jaú, Marília, Botucatu e Região"},
		"15": {"São Paulo": "Sorocaba e Região"},
		"16": {"São Paulo": "Ribeirão Preto, São Carlos, Araraquara e Região"},
		"17": {"São Paulo": "São José do Rio Preto e Região"},
		"18": {"São Paulo": "Presidente Prudente, Araçatuba e Região"},
		"19": {"São Paulo": "Região Metropolitana de Campinas"},
		"22": {"Rio de Janeiro": "Campos dos Goytacazes e Região"},
		"21": {"Rio de Janeiro": "Região Metropolitana do Rio de Janeiro"},
		"24": {"Rio de Janeiro": "Volta Redonda, Petrópolis e Região"},
		"27": {"Espírito Santo": "Região Metropolitana de Vitória"},
		"28": {"Espírito Santo": "Cachoeiro de Itapemirim e Região"},
		"31": {"Minas Gerais": "Região Metropolitana de Belo Horizonte"},
		"32": {"Minas Gerais": "Juiz de Fora e Região"},
		"33": {"Minas Gerais": "Governador Valadares e Região"},
		"34": {"Minas Gerais": "Uberlândia e região"},
		"35": {"Minas Gerais": "Poços de Caldas, Pouso Alegre, Varginha e Região"},
		"37": {"Minas Gerais": "Divinópolis, Itaúna e Região"},
		"38": {"Minas Gerais": "Montes Claros e Região"},
		// Sul
		"41": {"Paraná": "Região Metropolitana de Curitiba"},
		"42": {"Paraná": "Ponta Grossa e Região"},
		"43": {"Paraná": "Londrina e Região"},
		"44": {"Paraná": "Maringá e Região"},
		"45": {"Paraná": "Cascavel e Região"},
		"46": {"Paraná": "Francisco Beltrão, Pato Branco e Região"},
		"47": {"Santa Catarina": "Joinville, Blumenau, Balneário Camboriú e Região"},
		"48": {"Santa Catarina": "Região Metropolitana de Florianópolis e Criciúma"},
		"49": {"Santa Catarina": "Chapecó, Lages e Região"},
		"51": {"Rio Grande do Sul": "Região Metropolitana de Porto Alegre"},
		"53": {"Rio Grande do Sul": "Pelotas e Região"},
		"54": {"Rio Grande do Sul": "Caxias do Sul e Região"},
		"55": {"Rio Grande do Sul": "Santa Maria e Região"},
		// Centro-Oeste
		"61": {"Distrito Federal e Goiás": "Brasília e Região"},
		"62": {"Goiás": "Região Metropolitana de Goiânia"},
		"63": {"Tocantins": "Todos os municípios do estado"},
		"64": {"Goiás": "Rio Verde e Região"},
		"65": {"Mato Grosso": "Região Metropolitana de Cuiabá"},
		"66": {"Mato Grosso": "Rondonópolis e Região"},
		"67": {"Mato Grosso do Sul": "Todos os municípios do estado"},
		"68": {"Acre": "Todos os municípios do estado"},
		// Nordeste
		"71": {"Bahia": "Região Metropolitana de Salvador"},
		"73": {"Bahia": "Itabuna, Ilhéus e Região"},
		"74": {"Bahia": "Juazeiro e Região"},
		"75": {"Bahia": "Feira de Santana e Região"},
		"77": {"Bahia": "Vitória da Conquista e Região"},
		"79": {"Sergipe": "Todos os municípios do estado"},
		"81": {"Pernambuco": "Região Metropolitana de Recife"},
		"82": {"Alagoas": "Todos os municípios do estado"},
		"83": {"Paraíba": "Todos os municípios do estado"},
		"84": {"Rio Grande do Norte": "Todos os municípios do estado"},
		"85": {"Ceará": "Região Metropolitana de Fortaleza"},
		"86": {"Piauí": "Região de Teresina"},
		"87": {"Pernambuco": "Região de Petrolina"},
		"88": {"Ceará": "Região de Juazeiro do Norte"},
		"89": {"Piauí": "Região de Picos e Floriano"},
		// Norte
		"91": {"Pará": "Região Metropolitana de Belém"},
		"92": {"Amazonas": "Região de Manaus"},
		"93": {"Pará": "Região de Santarém"},
		"94": {"Pará": "Região de Marabá"},
		"95": {"Roraima": "Todos os municípios do estado"},
		"96": {"Amapá": "Todos os municípios do estado"},
		"97": {"Amazonas": "Região de Tefé e Coari"},
		"98": {"Maranhão": "Região Metropolitana de São Luís"},
		"99": {"Maranhão": "Região de Imperatriz"},
		"69": {"Rondônia": "Todos os municípios do estado"},
	}
)

type PhoneGen struct{}

func New() *PhoneGen {
	return &PhoneGen{}
}

func (p *PhoneGen) Random(limit int) []string {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	phones := []string{}

	// Generate and print random phone numbers
	for i := 0; i < limit; i++ {
		phones = append(phones, generateRandomPhoneNumber(rnd))
	}

	return phones
}

func (p *PhoneGen) RandomE164(limit int, countryCode string) ([]string, error) {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	phones := []string{}

	// Generate and print random phone numbers
	for i := 0; i < limit; i++ {
		phone := generateRandomPhoneNumber(rnd)
		e164, err := formatE164(phone, countryCode)
		if err != nil {
			return phones, err
		}
		phones = append(phones, e164)
	}

	return phones, nil
}

func (p *PhoneGen) RandomMobile(limit int) []string {
	// Create a new local random generator, seeded with Unix time.
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	phones := []string{}

	// Generate and print random phone numbers
	for i := 0; i < limit; i++ {
		// get a random area code
		ac := randomAreaCode(rnd)

		mobile := fmt.Sprintf("%s9%s", ac, randomDigits(rnd, 8))

		phones = append(phones, mobile)

	}

	return phones
}

func (p *PhoneGen) RandomLandline(limit int) []string {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	phones := []string{}

	// Generate and print random phone numbers
	for i := 0; i < limit; i++ {
		// get a random area code
		ac := randomAreaCode(rnd)

		// for landline, get a random number pattern
		pattern, err := getNumberPattern(ac)
		if err != nil {
			panic(err)
		}

		landline := fmt.Sprintf("%s%s%s", ac, pattern, randomDigits(rnd, 7))

		phones = append(phones, landline)

	}

	return phones
}

func (p *PhoneGen) RandomMobileE164(limit int, countryCode string) ([]string, error) {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	phones := []string{}

	// Generate and print random phone numbers
	for i := 0; i < limit; i++ {
		// get a random area code
		ac := randomAreaCode(rnd)

		mobile := fmt.Sprintf("%s9%s", ac, randomDigits(rnd, 8))

		e164, err := formatE164(mobile, countryCode)
		if err != nil {
			return phones, err
		}

		phones = append(phones, e164)

	}

	return phones, nil
}

func generateRandomPhoneNumber(rnd *rand.Rand) string {
	// get a random area code
	ac := randomAreaCode(rnd)

	// for landline, get a random number pattern
	pattern, err := getNumberPattern(ac)
	if err != nil {
		panic(err)
	}

	if rnd.Intn(2) == 0 {
		return fmt.Sprintf("%s%s%s", ac, pattern, randomDigits(rnd, 7))
	}

	return fmt.Sprintf("%s9%s", ac, randomDigits(rnd, 8))
}

func randomAreaCode(rnd *rand.Rand) string {
	ddd := rnd.Intn(len(areaCodes))

	for {
		_, ok := areaCodes[strconv.Itoa(ddd)]
		if !ok {
			ddd = rnd.Intn(len(areaCodes))
			continue
		}
		return strconv.Itoa(ddd)
	}
}

func randomDigits(rnd *rand.Rand, length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		fmt.Fprintf(&sb, "%d", rnd.Intn(10))
	}
	return sb.String()
}

func applyMask(phone string) string {
	// Remove all non-numeric characters
	regex := regexp.MustCompile(`\D`)
	cleanPhone := regex.ReplaceAllString(phone, "")

	// Apply mask based on length (landline has 10 digits, mobile has 11)
	if len(cleanPhone) == 10 {
		return fmt.Sprintf("(%s) %s-%s", cleanPhone[:2], cleanPhone[2:6], cleanPhone[6:])
	} else if len(cleanPhone) == 11 {
		return fmt.Sprintf("(%s) %s-%s", cleanPhone[:2], cleanPhone[2:7], cleanPhone[7:])
	}

	// If the phone number doesn't match expected lengths, just return
	return phone
}

func getNumberPattern(areaCode string) (string, error) {
	// For your information:
	// https://www.gov.br/anatel/pt-br/regulado/numeracao/tabela-servico-telefonico-fixo-comutado

	var pattern string

	switch areaCode {
	case "11", "12", "13", "14", "15", "16", "17", "18", "19":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "21", "22", "24":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "27", "28":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "31", "32", "33", "34", "35", "37", "38":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "41", "42", "43", "44", "45", "46":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "47", "48", "49":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "51", "53", "54", "55":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "61":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "62", "64":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "63":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "65", "66", "67", "68", "69":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "71", "73", "74", "75", "77":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "79":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "81", "87":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "82":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "83", "84":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "85", "88":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "86", "89":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "91", "93", "94":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "92", "97":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "95", "96":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	case "98", "99":
		pattern = "2XXX-XXXX; 3XXX-XXXX; 4XXX-XXXX; 5XXX-XXXX"
	default:
		return pattern, fmt.Errorf("unknown area code")
	}

	trimmed := strings.ReplaceAll(pattern, " ", "")
	pp := strings.Split(trimmed, ";")

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	r := pp[rnd.Intn(len(pp))]

	r = strings.ReplaceAll(r, "X", "")
	r = strings.ReplaceAll(r, "-", "")

	return r, nil
}

func formatE164(phone string, defaultCountryCode string) (string, error) {
	// Remove all non-numeric characters except the plus sign
	re := regexp.MustCompile(`[^\d+]`)
	phone = re.ReplaceAllString(phone, "")

	// Check if the phone number starts with '+'
	if !strings.HasPrefix(phone, "+") {
		phone = "+" + defaultCountryCode + phone
	}

	// Ensure the phone number does not exceed 15 digits
	if len(phone) > 16 { // + and 15 digits
		return "", fmt.Errorf("phone number exceeds the maximum length of 15 digits")
	}

	return phone, nil
}
