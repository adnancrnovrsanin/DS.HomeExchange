using Microsoft.AspNetCore.Identity;

namespace DS.HomeExchange.Auth.Models
{
    public class AppUser : IdentityUser
    {
        public string DisplayName { get; set; }
    }
}
