export const validateHexCode = (hex: string) => {
    return hex.match('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$')
}