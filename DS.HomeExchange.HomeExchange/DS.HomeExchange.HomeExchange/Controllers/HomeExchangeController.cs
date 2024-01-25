using DS.HomeExchange.HomeExchange.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Options;

namespace DS.HomeExchange.HomeExchange.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class HomeExchangeController : ControllerBase
    {
        private readonly AppDbContext _context;
        private readonly HttpClient _httpClient;
        private readonly IOptions<Urls> _urls;

        public HomeExchangeController(AppDbContext context, HttpClient httpClient, IOptions<Urls> urls)
        {
            _context = context;
            _httpClient = httpClient;
            _urls = urls;
        }

        [HttpGet]
        public async Task<ActionResult<HomeExchangeRequest>> Get()
        {
            var requests = await _context.HomeExhangeRequests.ToListAsync();

            return Ok(requests);
        }

        [HttpGet("{id}")]
        public async Task<ActionResult<HomeExchangeRequest>> Get(Guid id)
        {
            var request = await _context.HomeExhangeRequests.FindAsync(id);

            if (request == null)
            {
                return NotFound();
            }

            return Ok(request);
        }

        [HttpPost]
        public async Task<ActionResult<HomeExchangeRequest>> Post(HomeExchangeRequest request)
        {
            var newHomeExchangeRequest = new HomeExchangeRequest
            {
                FromUserId = request.FromUserId,
                ToUserId = request.ToUserId,
                FromUserHomeId = request.FromUserHomeId,
                ToUserHomeId = request.ToUserHomeId,
                Status = HomeExchangeStatus.Pending
            };

            _context.HomeExhangeRequests.Add(newHomeExchangeRequest);

            var result = await _context.SaveChangesAsync() > 0;

            if (!result)
            {
                return BadRequest();
            }

            return CreatedAtAction(nameof(Get), new { id = request.Id }, request);
        }

        [HttpPut]
        public async Task<ActionResult<HomeExchangeRequest>> Put(HomeExchangeRequest request)
        {
            _context.HomeExhangeRequests.Update(request);

            var result = await _context.SaveChangesAsync() > 0;
            if (!result)
            {
                return BadRequest();
            }

            return NoContent();
        }

        [HttpDelete("{id}")]
        public async Task<ActionResult<HomeExchangeRequest>> Delete(Guid id)
        {
            var request = await _context.HomeExhangeRequests.FindAsync(id);

            if (request == null)
            {
                return NotFound();
            }

            _context.HomeExhangeRequests.Remove(request);

            var result = await _context.SaveChangesAsync() > 0;
            if (!result)
            {
                return BadRequest();
            }

            return NoContent();
        }

        [HttpPut("{id}/accept")]
        public async Task<ActionResult<HomeExchangeRequest>> Accept(Guid id)
        {
            var request = await _context.HomeExhangeRequests.FindAsync(id);

            if (request == null)
            {
                return NotFound();
            }

            var response = await _httpClient.PostAsJsonAsync($"{_urls.Value.Homes}/{request.ToUserHomeId}/accept", request);
            if (!response.IsSuccessStatusCode)
            {
                return BadRequest("Error when trying to accept the exchange.");
            }

            request.Status = HomeExchangeStatus.Accepted;

            _context.HomeExhangeRequests.Update(request);

            var result = await _context.SaveChangesAsync() > 0;
            if (!result)
            {
                return BadRequest();
            }

            return NoContent();
        }

        [HttpPut("{id}/reject")]
        public async Task<ActionResult<HomeExchangeRequest>> Reject(Guid id)
        {
            var request = await _context.HomeExhangeRequests.FindAsync(id);

            if (request == null)
            {
                return NotFound();
            }

            HttpResponseMessage response = await _httpClient.PostAsJsonAsync($"{_urls.Value.Homes}/{request.ToUserHomeId}/reject", "");
            if (!response.IsSuccessStatusCode)
            {
                return BadRequest("Error when trying to reject the exchange.");
            }

            request.Status = HomeExchangeStatus.Rejected;

            _context.HomeExhangeRequests.Update(request);

            var result = await _context.SaveChangesAsync() > 0;
            if (!result)
            {
                return BadRequest();
            }

            return NoContent();
        }
    }
}
