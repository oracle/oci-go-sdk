import xml.etree.ElementTree as ET
import pytest
from add_or_update_scripts.python_sdk_add_or_update_spec import GENERATE_EXECUTION_TEMPLATE


def test_format_general_execution_template():

    spec_name = "datasafe"
    spec_generation_type = "PREVIEW"
    artifact_id = "ads-control-plane-spec"
    pom_var_code_gen_language = "${codegen-language}"
    pom_var_preprocessed_temp_dir = "${preprocessed-temp-dir}"
    pom_var_feature_id_file = "${feature-id-file}"
    pom_var_feature_id_dir = "${feature-id-dir}"

    result = GENERATE_EXECUTION_TEMPLATE.format(
        artifact_id=artifact_id,
        spec_name=spec_name,
        spec_generation_type=spec_generation_type,
        regional_non_regional_service_overrides=""
    )

    assert(artifact_id in result)
    assert(spec_name in result)
    assert(spec_generation_type in result)
    assert(pom_var_code_gen_language in result)
    assert(pom_var_preprocessed_temp_dir in result)
    assert(pom_var_feature_id_file in result)
    assert(pom_var_feature_id_dir in result)
    assert("isRegionalClient" not in result)
    print(result)
    ET.fromstring(result)
    try:
        ET.fromstring(result)
    except Exception as e:
        pytest.fail("Unexpected error parsing XML string")
        print(e)

    non_regional_service_overide = '<isRegionalClient.{service_name}>false</isRegionalClient.{service_name}>\n'.format(service_name="datasafe")
    result = result = GENERATE_EXECUTION_TEMPLATE.format(
        artifact_id=artifact_id,
        spec_name=spec_name,
        spec_generation_type=spec_generation_type,
        regional_non_regional_service_overrides=non_regional_service_overide
    )

    assert("isRegionalClient" in result)
    try:
        ET.fromstring(result)
    except Exception as e:
        pytest.fail("Unexpected error parsing XML string")
        print(e)
