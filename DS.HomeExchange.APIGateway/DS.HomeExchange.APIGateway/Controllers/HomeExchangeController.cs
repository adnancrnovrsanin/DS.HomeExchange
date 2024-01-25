using DS.HomeExchange.APIGateway.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Options;

namespace DS.HomeExchange.APIGateway.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class HomeExchangeController : ControllerBase
    {
        private readonly HttpClient _httpClient;
        private readonly IOptions<Urls> _urls;
        public HomeExchangeController(HttpClient httpClient, IOptions<Urls> urls)
        {
            _httpClient = httpClient;
            _urls = urls;
        }

        [HttpGet]
        public async Task<IActionResult> Get()
        {
            var response = await _httpClient.GetAsync($"{_urls.Value.HomeExchange}/api/HomeExchange");
            if (response.IsSuccessStatusCode)
            {
                var content = await response.Content.ReadAsStringAsync();
                return Ok(content);
            }
            return NotFound();
        }

        [HttpGet("{id}")]
        public async Task<IActionResult> Get(string id)
        {
            var response = await _httpClient.GetAsync($"{_urls.Value.HomeExchange}/api/HomeExchange/{id}");
            if (response.IsSuccessStatusCode)
            {
                var content = await response.Content.ReadAsStringAsync();
                return Ok(content);
            }
            return NotFound();
        }

        [HttpPost]
        public async Task<IActionResult> Post([FromBody] HomeExchangeRequest homeExchange)
        {
            var response = await _httpClient.PostAsJsonAsync($"{_urls.Value.HomeExchange}/api/HomeExchange", homeExchange);
            if (response.IsSuccessStatusCode)
            {
                var content = await response.Content.ReadAsStringAsync();
                return Ok(content);
            }
            return NotFound();
        }

        [HttpPut]
        public async Task<IActionResult> Put([FromBody] HomeExchangeRequest homeExchange)
        {
            var response = await _httpClient.PutAsJsonAsync($"{_urls.Value.HomeExchange}/api/HomeExchange", homeExchange);
            if (response.IsSuccessStatusCode)
            {
                var content = await response.Content.ReadAsStringAsync();
                return Ok(content);
            }
            return NotFound();
        }

        [HttpDelete("{id}")]
        public async Task<IActionResult> Delete(string id)
        {
            var response = await _httpClient.DeleteAsync($"{_urls.Value.HomeExchange}/api/HomeExchange/{id}");
            if (response.IsSuccessStatusCode)
            {
                var content = await response.Content.ReadAsStringAsync();
                return Ok(content);
            }
            return NotFound();
        }

        [HttpPut("{id}/accept")]
        public async Task<IActionResult> Accept(string id)
        {
            var response = await _httpClient.PutAsync($"{_urls.Value.HomeExchange}/api/HomeExchange/{id}/accept", null);
            if (response.IsSuccessStatusCode)
            {
                var content = await response.Content.ReadAsStringAsync();
                return Ok(content);
            }
            return NotFound();
        }

        [HttpPut("{id}/reject")]
        public async Task<IActionResult> Reject(string id)
        {
            var response = await _httpClient.PutAsync($"{_urls.Value.HomeExchange}/api/HomeExchange/{id}/reject", null);
            if (response.IsSuccessStatusCode)
            {
                var content = await response.Content.ReadAsStringAsync();
                return Ok(content);
            }
            return NotFound();
        }
    }
}
