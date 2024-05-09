package controller

import (
    "github.com/aiteung/musik"
    "github.com/gofiber/fiber/v2"

	cek "github.com/gryzlegrizz/package_uts/module"
)

func Homepage(c *fiber.Ctx) error {
    ipaddr := musik.GetIPaddress()
    return c.JSON(ipaddr)
}

func GetAllMuseumCollections(c *fiber.Ctx) error {
	ps := cek.GetAllMuseumCollections()
	return c.JSON(ps)
	}