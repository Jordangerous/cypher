-- SQL schema for the sample database
-- Definition of the main table 'Person' with various data types
CREATE TABLE
    Person (
        id VARCHAR(36) PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        age INT,
        email VARCHAR(255),
        isActive BOOLEAN,
        createdAt DATETIME,
        address_id INT,
        FOREIGN KEY (address_id) REFERENCES Address (id)
    );

-- Definition of the 'Address' table for the N:1 relationship
CREATE TABLE
    Address (
        id INT PRIMARY KEY AUTO_INCREMENT,
        street VARCHAR(255),
        city VARCHAR(100),
        postalCode VARCHAR(20)
    );

-- Definition of the 'Group' table representing a 1:N relationship
CREATE TABLE
    Group (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(100)
    );

-- Definition of the 'Organization' table representing the N:N relationship
CREATE TABLE
    Organization (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(100)
    );

-- Join table for the 1:N relationship between 'Group' and 'Person'
CREATE TABLE
    GroupMembers (
        group_id INT,
        person_id VARCHAR(36),
        PRIMARY KEY (group_id, person_id),
        FOREIGN KEY (group_id) REFERENCES Group (id),
        FOREIGN KEY (person_id) REFERENCES Person (id)
    );

-- Join table for the N:N relationship between 'Organization' and 'Person'
CREATE TABLE
    OrganizationEmployees (
        organization_id INT,
        person_id VARCHAR(36),
        PRIMARY KEY (organization_id, person_id),
        FOREIGN KEY (organization_id) REFERENCES Organization (id),
        FOREIGN KEY (person_id) REFERENCES Person (id)
    );