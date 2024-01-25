#See https://aka.ms/customizecontainer to learn how to customize your debug container and how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
USER app
WORKDIR /app
EXPOSE 8080
EXPOSE 8081

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src
COPY ["DS.HomeExchange.Auth/DS.HomeExchange.Auth.csproj", "DS.HomeExchange.Auth/"]
RUN dotnet restore "./DS.HomeExchange.Auth/./DS.HomeExchange.Auth.csproj"
COPY . .
WORKDIR "/src/DS.HomeExchange.Auth"
RUN dotnet build "./DS.HomeExchange.Auth.csproj" -c $BUILD_CONFIGURATION -o /app/build

FROM build AS publish
ARG BUILD_CONFIGURATION=Release
RUN dotnet publish "./DS.HomeExchange.Auth.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "DS.HomeExchange.Auth.dll"]