using static System.Runtime.InteropServices.JavaScript.JSType;

namespace DS.HomeExchange.APIGateway.DTOs
{
    public class GoAPIResponse
    {
        public int status { get; set; }
        public string message { get; set; }
        public Data data { get; set; }
    }

    public class Data
    {
        public List<HomeGoAPIDto> data { get; set; }
    }

    public class HomeGoAPIDto
    {
        public string id { get; set; }
        public string ownerId { get; set; }
        public string address { get; set; }
        public string description { get; set; }
        public string createdAt { get; set; }
    }
}
