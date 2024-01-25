using AutoMapper;
using DS.HomeExchange.APIGateway.DTOs;
using DS.HomeExchange.APIGateway.Models;
using DS.HomeExchange.APIGateway.Services;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Options;

namespace DS.HomeExchange.APIGateway.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class HomeController : ControllerBase
    {
        private readonly HttpClient _httpClient;
        private readonly IOptions<Urls> _urls;

        public HomeController(HttpClient httpClient, IOptions<Urls> urls)
        {
            _httpClient = httpClient;
            _urls = urls;
        }

        [HttpGet]
        public async Task<IActionResult> Get()
        {
            var response = await _httpClient.GetAsync($"{_urls.Value.Homes}/");
            if (response.IsSuccessStatusCode)
            {
                var result = CustomMappingService.MapGoAPIResponseToHome(await response.Content.ReadFromJsonAsync<GoAPIResponse>());
                return Ok(result);
            }

            return BadRequest();
        }

        [HttpGet("{id}")]
        public async Task<IActionResult> Get(string id)
        {
            var response = await _httpClient.GetAsync($"{_urls.Value.Homes}/{id}");
            if (response.IsSuccessStatusCode)
            {
                var result = CustomMappingService.MapGoAPIResponseToHome(await response.Content.ReadFromJsonAsync<GoAPIResponse>());
                return Ok(result);
            }

            return BadRequest();
        }

        [HttpGet("owner/{id}")]
        public async Task<IActionResult> GetByOwner(string id)
        {
            var response = await _httpClient.GetAsync($"{_urls.Value.Homes}/owner/{id}");
            if (response.IsSuccessStatusCode)
            {
                var result = CustomMappingService.MapGoAPIResponseToHome(await response.Content.ReadFromJsonAsync<GoAPIResponse>());
                return Ok(result);
            }

            return BadRequest();
        }

        [HttpPost]
        public async Task<IActionResult> Post(Home request)
        {
            var response = await _httpClient.PostAsJsonAsync($"{_urls.Value.Homes}/", request);
            if (response.IsSuccessStatusCode)
            {
                var result = CustomMappingService.MapGoAPIResponseToHome(await response.Content.ReadFromJsonAsync<GoAPIResponse>());
                return Ok(result);
            }

            return BadRequest();
        }

        [HttpPut]
        public async Task<IActionResult> Put(Home request)
        {
            var response = await _httpClient.PutAsJsonAsync($"{_urls.Value.Homes}/", request);
            if (response.IsSuccessStatusCode)
            {
                var result = CustomMappingService.MapGoAPIResponseToHome(await response.Content.ReadFromJsonAsync<GoAPIResponse>());
                return Ok(result);
            }

            return BadRequest();
        }

        [HttpDelete("{id}")]
        public async Task<IActionResult> Delete(string id)
        {
            var response = await _httpClient.DeleteAsync($"{_urls.Value.Homes}/{id}");
            if (response.IsSuccessStatusCode)
            {
                var result = CustomMappingService.MapGoAPIResponseToHome(await response.Content.ReadFromJsonAsync<GoAPIResponse>());
                return Ok(result);
            }

            return BadRequest();
        }
    }
}
