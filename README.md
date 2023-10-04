# XLSXtoMySQL-Importer

Bu araç, Excel (XLSX) dosyalarındaki verileri hızlı ve kolay bir şekilde MySQL veritabanına aktarmak için tasarlanmıştır. Veri aktarımını kolaylaştırmak ve hızlandırmak için özelleştirilebilir konfigürasyon seçenekleri sunar.
- Bu projeyi geliştirmek veya hataları düzeltmek isterseniz, lütfen GitHub deposuna katkıda bulunun.
## Kullanım

1. **Konfigürasyon**: `config.toml` dosyasını düzenleyerek MySQL veritabanı bilgilerinizi ve hedef tablo adınızı ayarlayın.
2. **MySQL Tablo Yapısının Hazırlığı**: İşlem yapmak istediğiniz Excel dosyasındaki sütunların isimleri ile bire bir aynı isimde bir MySQL tablosu oluşturmanız gerekmektedir. Bu adım, Excel dosyasındaki verilerin düzgün bir şekilde veritabanına aktarılabilmesi için önemlidir.

   Örnek Excel Tablosu:

   | Num  | FirstName | LastName | Email             | Salary |
   | --- | --------- | -------- | ----------------- | ------ |
   | 1   | Ahmet     | Kaya     | ahmet@example.com | 50000  |
   | 22   | Elif      | Tekin    | elif@example.com  | 55000  |
   | 33   | Mehmet    | Şahin    | mehmet@example.com| 60000  |

   Örnek SQL Sorgusu:

   Aşağıdaki SQL sorgusu, yukarıdaki örnek Excel tablosunun karşılığı olan bir MySQL tablosunu oluşturur:

   ```sql
   CREATE TABLE sample_table (
       ID INT AUTO_INCREMENT PRIMARY KEY,
       Num INT(11) NULL DEFAULT NULL,
       FirstName VARCHAR(255),
       LastName VARCHAR(255),
       Email VARCHAR(255),
       Salary INT
   );

# XLSXtoMySQL-Importer

This tool is designed to facilitate the quick and easy transfer of data from Excel (XLSX) files to a MySQL database. It offers customizable configuration options to streamline and accelerate the data transfer process.

- If you'd like to contribute to or fix issues in this project, please consider contributing to our

## Usage

1. **Configuration**: Edit the `config.toml` file to set your MySQL database information and target table name.
2. **Preparation of MySQL Table Structure**: You need to create a MySQL table with column names identical to those in the Excel file you want to process. This step is crucial to ensure that the data from the Excel file is transferred accurately to the database.

   Sample Excel Table:

   | Num  | FirstName | LastName | Email             | Salary |
   | --- | --------- | -------- | ----------------- | ------ |
   | 1   | Ahmet     | Kaya     | ahmet@example.com | 50000  |
   | 22   | Elif     | Tekin    | elif@example.com  | 55000  |
   | 33   | Mehmet    | Lale    | mehmet@example.com| 60000  |

   Sample SQL Query:

   The following SQL query creates a MySQL table that corresponds to the sample Excel table above:

   ```sql
   CREATE TABLE sample_table (
       ID INT AUTO_INCREMENT PRIMARY KEY,
       Num INT(11) NULL DEFAULT NULL,
       FirstName VARCHAR(255),
       LastName VARCHAR(255),
       Email VARCHAR(255),
       Salary INT
   );

