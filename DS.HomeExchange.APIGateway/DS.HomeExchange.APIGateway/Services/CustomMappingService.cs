using DS.HomeExchange.APIGateway.DTOs;
using DS.HomeExchange.APIGateway.Models;

namespace DS.HomeExchange.APIGateway.Services
{
    public static class CustomMappingService
    {
        public static List<Home> MapGoAPIResponseToHome(GoAPIResponse source)
        {
            if (source == null) return null;
            if (source.data == null) return null;
            if (source.data.data == null) return null;

            List<Home> result = [];


            foreach (HomeGoAPIDto homeDto in source.data.data)
            {
                Home newHome = new()
                {
                    Id = homeDto.id,
                    OwnerId = homeDto.ownerId,
                    Address = homeDto.address,
                    Description = homeDto.description,
                    CreatedAt = homeDto.createdAt
                };
                result.Add(newHome);
            }

            return result;
        }

        public static HomeGoAPIDto MapHomeToHomeGoAPIDto(Home home)
        {
            if (home == null) return null;

            return new()
            {
                id = home.Id,
                ownerId = home.OwnerId,
                description = home.Description,
                address = home.Address,
                createdAt = home.CreatedAt
            };  
        }
    }
}
