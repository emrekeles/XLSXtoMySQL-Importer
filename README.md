# XLSXtoMySQL-Importer

Bu araç, Excel (XLSX) dosyalarındaki verileri hızlı ve kolay bir şekilde MySQL veritabanına aktarmak için tasarlanmıştır. Veri aktarımını kolaylaştırmak ve hızlandırmak için özelleştirilebilir konfigürasyon seçenekleri sunar.
- Bu projeyi geliştirmek veya hataları düzeltmek isterseniz, lütfen GitHub deposuna katkıda bulunun.
## Kullanım

1. **Konfigürasyon**: `config.toml` dosyasını düzenleyerek MySQL veritabanı bilgilerinizi ve hedef tablo adınızı ayarlayın.
2. **MySQL Tablo Yapısının Hazırlığı**: İşlem yapmak istediğiniz Excel dosyasındaki sütunların isimleri ile bire bir aynı isimde bir MySQL tablosu oluşturmanız gerekmektedir. Bu adım, Excel dosyasındaki verilerin düzgün bir şekilde veritabanına aktarılabilmesi için önemlidir.

   Örnek Excel Tablosu:

   | ID  | FirstName | LastName | Email             | Salary |
   | --- | --------- | -------- | ----------------- | ------ |
   | 1   | Ahmet     | Kaya     | ahmet@example.com | 50000  |
   | 2   | Ayşe      | Tekin    | ayse@example.com  | 55000  |
   | 3   | Mehmet    | Şahin    | mehmet@example.com| 60000  |

   Örnek SQL Sorgusu:

   Aşağıdaki SQL sorgusu, yukarıdaki örnek Excel tablosunun karşılığı olan bir MySQL tablosunu oluşturur:

   ```sql
   CREATE TABLE employees (
       ID INT AUTO_INCREMENT PRIMARY KEY,
       FirstName VARCHAR(255),
       LastName VARCHAR(255),
       Email VARCHAR(255),
       Salary INT
   );


