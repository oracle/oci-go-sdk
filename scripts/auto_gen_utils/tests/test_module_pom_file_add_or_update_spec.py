import xml.etree.ElementTree as ET
from add_or_update_scripts.module_pom_file_add_or_update_spec import add_module_to_parent_pom
from add_or_update_scripts.python_sdk_add_or_update_spec import MODULE_TEMPLATE


pom_string = """<?xml version='1.0' encoding='UTF-8'?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
  <modules>
    <module>poms/budget/pom.xml</module>
    <module>poms/core/pom.xml</module>
  </modules>
</project>
"""

ns = {"ns": "http://maven.apache.org/POM/4.0.0"}


def test_add_module_to_parent_pom():
    spec_name = "datasafe"
    module = MODULE_TEMPLATE.format(spec_name)
    xpath = ".//ns:modules"
    pom = ET.fromstring(pom_string)
    properties = pom.findall(xpath, ns)[0]
    starting_length = len(properties)

    add_module_to_parent_pom(pom, module)
    properties = pom.findall(xpath, ns)[0]
    ending_length = len(properties)

    assert(ending_length > starting_length)
    # The first module should be the one just added.
    assert(spec_name in properties[0].text)
