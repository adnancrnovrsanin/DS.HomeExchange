using DS.HomeExchange.HomeExchange.Models;
using Microsoft.EntityFrameworkCore;

namespace DS.HomeExchange.HomeExchange
{
    public class AppDbContext : DbContext
    {
        public AppDbContext(DbContextOptions options) : base(options)
        {
        }

        protected AppDbContext()
        {
        }

        public DbSet<HomeExchangeRequest> HomeExhangeRequests { get; set; }
    }
}
