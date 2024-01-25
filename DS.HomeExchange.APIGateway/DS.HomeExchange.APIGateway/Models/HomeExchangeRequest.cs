namespace DS.HomeExchange.APIGateway.Models
{
    public class HomeExchangeRequest
    {
        public Guid Id { get; set; }
        public string FromUserId { get; set; }
        public string ToUserId { get; set; }
        public string FromUserHomeId { get; set; }
        public string ToUserHomeId { get; set; }
        public HomeExchangeStatus Status { get; set; }
    }
}
